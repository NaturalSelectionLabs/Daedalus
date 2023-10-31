import * as yaml from "js-yaml";
import * as core from "@actions/core";

interface App {
  name: string;
  namespace: string;
  cluster: string;
  project: string;
  repo: string;
  revision: string;
  image: Image;
  sync: boolean;
  directory: string;
}

interface Image {
  name: string;
  tag: string;
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
    directory: core.getInput("dir"),
    image: {
      name: core.getInput("image-name"),
      tag: core.getInput("image-tag"),
    },
    sync: core.getBooleanInput("auto-sync"),
  };
};

export function build(a: App) {
  const plugin: any = {
    repoURL: `https://github.com/${a.repo}`,
    targetRevision: a.revision,
    path: a.directory,
    plugin: {
      name: "avp-kustomize",
      env: [
        {
          name: "AVP_SECRET",
          value: `guardian:avp-${a.cluster}`,
        },
        {
          name: "APP_REPO",
          value: a.repo,
        },
        {
          name: "IMAGE_NAME",
          value: a.image.name,
        },
        {
          name: "IMAGE_TAG",
          value: a.image.tag,
        },
      ],
    },
  };

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
      source: plugin,
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
