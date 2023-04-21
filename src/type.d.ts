interface ModelProps<T> {
  value: T;
  onChange: (value: T) => void;
}

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
