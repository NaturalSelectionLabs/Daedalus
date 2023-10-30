interface ModelProps<T> {
  value: T;
  onChange: (value: T) => void;
}

type NestedPartial<T> = {
  [K in keyof T]?: T[K] extends object ? NestedPartial<T[K]> : T[K];
};

interface TriggerProps {
  push: {
    branches: string[];
    tags: string[];
  };
}

type JobProps = Partial<{
  name: string;
  runsOn: string;
  steps: StepProps[];
  needs: string[];
  if: string;
  env: {};
  uses: string;
  with: {};
  secrets: "inherit" | {};
}>;

interface StepProps {}

interface GithubActionProps {
  name: string;
  on: TriggerProps;
  jobs: {
    [key: string]: JobProps;
  };
}

type KVMap = map<string, string>;

interface PodResource {
  cpu: string;
  memory: string;
}

interface HelmValues {
  replicaCount: number;
  image: {
    repository: string;
    pullPolicy: string;
    tag: string;
  };
  imagePullSecrets: string[];
  nameOverride: string;
  fullnameOverride: string;
  env: KVMap;
  envFrom: string[];
  serviceAccount: {
    create: boolean;
    annotations: KVMap;
    name: string;
  };
  podAnnotations: KVMap;
  podSecurityContext: {
    fsGroup: number;
  };
  securityContext: {
    runAsUser: number;
    runAsNonRoot: boolean;
    readOnlyRootFilesystem: boolean;
    capabilities: {
      drop: string[];
    };
  };
  service: {
    type: string;
    port: number;
  };
  ingressRoute: {
    enabled: boolean;
    annotations: KVMap;
    entrypoint: string[];
    routes: string[];
    tls: {
      secretName: string;
    };
  };
  resources: {
    limits: PodResource;
    requests: PodResource;
  };
  autoscaling: {
    enabled: boolean;
    minReplicas: number;
    maxReplicas: number;
    targetCPUUtilizationPercentage: number;
  };
  nodeSelector: KVMap;
  tolerations: any[];
  affinity: any;
}
