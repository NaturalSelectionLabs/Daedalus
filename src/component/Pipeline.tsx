import Container from "@mui/material/Container";
import Trigger from "./Trigger";
import BuildDocker from "./BuildDocker";
import Deploy from "./Deploy";
import React from "react";

const clusters = ["dev", "prod"];

const Pipeline = (props: ModelProps<GithubActionProps>) => {
  const setTrigger = (trigger: TriggerProps) => {
    props.onChange({
      ...props.value,
      on: trigger,
    });
  };

  const buildJob = props.value.jobs["build-and-push"];

  const setJob = (key: string, job: JobProps | undefined) => {
    const jobs = props.value.jobs;

    if (job) {
      jobs[key] = job;
      props.onChange({
        ...props.value,
        jobs,
      });
    } else {
      const newJobs = { ...jobs };
      delete newJobs[key];
      props.onChange({
        ...props.value,
        jobs: newJobs,
      });
    }
  };

  return (
    <Container style={{ margin: "20px 0 0 20px" }} maxWidth="sm">
      <Trigger value={props.value.on} onChange={setTrigger} />
      <div style={{ marginTop: "20px" }}></div>
      <BuildDocker
        value={buildJob}
        onChange={(val) => setJob("build-and-push", val)}
      />
      {clusters.map((cluster) => {
        return (
          <React.Fragment>
            <div style={{ marginTop: "20px" }}></div>
            <Deploy
              value={props.value.jobs[`deploy-to-${cluster}`]}
              onChange={(val) => setJob(`deploy-to-${cluster}`, val)}
              env={cluster}
            />
          </React.Fragment>
        );
      })}
    </Container>
  );
};

export default Pipeline;
