import * as yaml from "js-yaml";
import * as core from "@actions/core";

interface App {
  name: string;
  namespace: string;
  cluster: string;
  project: string;
  repo: string;
  revision: string;
  helm: HelmApp;
  kustomize: KustomizeApp;
  image: Image;
  sync: boolean;
}

interface Image {
  name: string;
  tag: string;
}

interface HelmApp {
  valueFiles: string[];
  releaseName: string;
  chart: HelmChart;
}

interface HelmChart {
  repoUrl: string;
  name: string;
  version: string;
}

interface KustomizeApp {
  directory: string;
}

const valueFilePathRemovePrefix = (inputPath: string): string => {
  if (inputPath.startsWith("./")) {
    return inputPath.slice(2);
  } else if (inputPath.startsWith("/")) {
    return inputPath.slice(1);
  } else {
    return inputPath;
  }
};

export const load = (): App => {
  const name = core.getInput("name");
  const namespace = core.getInput("namespace");
  const cluster = core.getInput("cluster");
  const project = core.getInput("project") || namespace;
  const repo = core.getInput("repo");
  const revision = core.getInput("revision");

  return {
    name,
    namespace,
    cluster,
    project,
    repo,
    revision,
    helm: {
      valueFiles: core.getMultilineInput("helm-value-files"),
      releaseName: core.getInput("helm-release-name") || name,
      chart: {
        repoUrl: core.getInput("helm-chart-url"),
        name: core.getInput("helm-chart-name"),
        version: core.getInput("helm-chart-version"),
      },
    },
    kustomize: {
      directory: core.getInput("kustomize-dir"),
    },
    image: {
      name: core.getInput("image-name"),
      tag: core.getInput("image-tag"),
    },
    sync: core.getBooleanInput("auto-sync"),
  };
};

export function build(a: App) {
  const ref: any = {
    ref: "values",
    repoURL: `https://github.com/${a.repo}`,
    targetRevision: a.revision,
  };

  const helm: any = {
    chart: a.helm.chart.name,
    helm: {
      releaseName: a.helm.releaseName,
      valueFiles: a.helm.valueFiles.map(
        (item) => `\\\$values/${valueFilePathRemovePrefix(item)}`,
      ),
      parameters: [
        { name: "image.tag", value: a.image.tag },
        { name: "repoUrl", value: a.repo },
      ],
    },
    repoURL: a.helm.chart.repoUrl,
    targetRevision: a.helm.chart.version,
  };

  const kustomize: any = {
    repoURL: `https://github.com/${a.repo}`,
    targetRevision: a.revision,
    path: a.kustomize.directory,
    kustomize: {
      images: [`${a.image.name}:${a.image.tag}`],
      commonAnnotations: {
        "github.com/url": a.repo,
      },
    },
  };

  const applicationSources = [];

  if (a.helm.valueFiles.length > 0) {
    applicationSources.push(helm);
    applicationSources.push(ref);
  }

  if (a.kustomize.directory !== "") {
    applicationSources.push(kustomize);
  }

  const application = {
    apiVersion: "argoproj.io/v1alpha1",
    kind: "Application",
    metadata: {
      name: `${a.name}-${a.cluster}`,
      namespace: "guardian",
    },
    spec: {
      destination: {
        namespace: a.namespace,
        name: a.cluster,
      },
      project: a.project,
      sources: applicationSources,
      syncPolicy: {
        syncOptions: ["ApplyOutOfSyncOnly=true", "ServerSideApply=true"],
        automated: a.sync ? {} : undefined,
      },
    },
  };

  return application;
}

export function toYaml(a: App): string {
  const jsonData = JSON.stringify(build(a), null, 2);
  const data = JSON.parse(jsonData);

  // delete data.metadata.creationTimestamp;
  // delete data.status;

  return yaml.dump(data);
}

function imageNameString(i: Image): string {
  return `${i.name}:${i.tag}`;
}
