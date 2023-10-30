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

function valueFilePathRemovePrefix(inputPath: string): string {
  if (inputPath.startsWith("./")) {
    return inputPath.slice(2);
  } else {
    return inputPath;
  }
}

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
  };
};

export function build(a: App) {
  const ref: any = {
    Ref: "values",
    RepoURL: a.repo,
    TargetRevision: a.revision,
  };

  const helm: any = {
    Chart: a.helm.chart.name,
    Helm: {
      ReleaseName: a.helm.releaseName,
      ValueFiles: a.helm.valueFiles.map(
        (item) => `$values/${valueFilePathRemovePrefix(item)}`,
      ),
      Parameters: [
        { Name: "image.tag", Value: a.image.tag },
        { Name: "repoUrl", Value: a.repo },
      ],
    },
    RepoURL: a.helm.chart.repoUrl,
    TargetRevision: a.helm.chart.version,
  };

  const kustomize: any = {
    RepoURL: a.repo,
    TargetRevision: a.revision,
    Path: a.kustomize.directory,
    Kustomize: {
      Images: [`${a.image.name}:${a.image.tag}`],
      CommonAnnotations: {
        "github.com/url": a.repo,
      },
    },
  };

  const applicationSources = [];

  if (a.helm.valueFiles.length > 0) {
    applicationSources.push(helm);
  }

  if (a.kustomize.directory !== "") {
    applicationSources.push(kustomize);
  }

  applicationSources.push(ref);

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
    },
  };

  return application;
}

export function toYaml(a: App): string {
  const jsonData = JSON.stringify(build(a), null, 2);
  const data = JSON.parse(jsonData);
  console.log(yaml.dump(data));

  // delete data.metadata.creationTimestamp;
  // delete data.status;

  return yaml.dump(data);
}

function imageNameString(i: Image): string {
  return `${i.name}:${i.tag}`;
}
