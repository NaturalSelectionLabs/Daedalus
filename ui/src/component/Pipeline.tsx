import Container from "@mui/material/Container";
import Trigger from "./Trigger";
import BuildDocker from "./BuildDocker";
import Deploy from "./Deploy";
import React, { useState } from "react";
import {
  Card,
  CardContent,
  FormControl,
  InputLabel,
  MenuItem,
  Select,
  SelectChangeEvent,
  TextField,
} from "@mui/material";

const clusters = ["dev", "prod"];

const namespaces = ["default", "pregod", "guardian", "tureco"];
const charts = ["web-app", "test"];

const Pipeline = (props: ModelProps<GithubActionProps>) => {
  const [releaseName, setReleaseName] = useState<string>("appName");

  const setTrigger = (trigger: TriggerProps) => {
    props.onChange({
      ...props.value,
      on: trigger,
    });
  };

  const buildJob = props.value.jobs["build"];

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

  const [namespace, setNamespace] = useState<string>("default");
  const [chart, setChart] = useState<string>("web-app");
  const [images, setImages] = useState<string>("repo/imageName");
  const [release, setRelease] = useState<string>("exampleApp");

  const setJobFields = (j: string, k: string, v: string) => {
    if (!props.value.jobs[j]) return;
    setJob(j, {
      ...props.value.jobs[j],
      with: {
        ...props.value.jobs[j]?.with,
        [k]: v,
      },
    });
  };

  const handleChangeImages = (event: React.ChangeEvent<HTMLInputElement>) => {
    const i = event.target.value;
    setImages(i);
    // set dev and prod images
    setJobFields("build", "images", i);
    setJobFields("deploy-dev", "images", i);
    setJobFields("deploy-prod", "images", i);
  };

  const handleChangeNamespace = (event: SelectChangeEvent) => {
    const n = event.target.value as string;
    setNamespace(n);
    // set dev and prod namespace
    setJobFields("deploy-dev", "namespace", n);
    setJobFields("deploy-prod", "namespace", n);
  };

  const handleChangeChart = (event: SelectChangeEvent) => {
    const c = event.target.value as string;
    setChart(c);
    // set dev and prod chart
    setJobFields("deploy-dev", "chart", c);
    setJobFields("deploy-prod", "chart", c);
  };

  const handleChangeReleaseName = (
    event: React.ChangeEvent<HTMLInputElement>
  ) => {
    const r = event.target.value;
    setReleaseName(r);
    // set dev and prod releaseName
    setJobFields("deploy-dev", "releaseName", r);
    setJobFields("deploy-prod", "releaseName", r);
    props.onChange({
      ...props.value,
      name: `Build and Deploy ${r} Workflow`,
    });
  };

  return (
    <Container style={{ margin: "20px 0 0 20px" }} maxWidth="sm">
      <Trigger value={props.value.on} onChange={setTrigger} />
      <div style={{ marginTop: "20px" }}></div>
      <Card>
        <CardContent>
          <FormControl fullWidth>
            <TextField
              label="Image"
              variant="outlined"
              value={images}
              onChange={handleChangeImages}
            />
            <div style={{ marginTop: "20px" }}></div>
            <TextField
              label="App"
              variant="outlined"
              value={releaseName}
              onChange={handleChangeReleaseName}
            />
          </FormControl>
        </CardContent>
      </Card>
      <BuildDocker
        value={buildJob}
        images={images}
        onChange={(val) => setJob("build", val)}
      />
      <Card>
        <CardContent>
          <FormControl fullWidth>
            <InputLabel>Namespace</InputLabel>
            <Select
              value={namespace}
              label="Namespace"
              onChange={handleChangeNamespace}
            >
              {namespaces.map((n) => (
                <MenuItem key={n} value={n}>
                  {n}
                </MenuItem>
              ))}
            </Select>
          </FormControl>
          <FormControl fullWidth style={{ marginTop: "10px" }}>
            <InputLabel>Chart</InputLabel>
            <Select value={chart} label="Chart" onChange={handleChangeChart}>
              {charts.map((n) => (
                <MenuItem key={n} value={n}>
                  {n}
                </MenuItem>
              ))}
            </Select>
          </FormControl>
          {clusters.map((cluster) => {
            return (
              <React.Fragment>
                <div style={{ marginTop: "20px" }}></div>
                <Deploy
                  value={props.value.jobs[`deploy-${cluster}`]}
                  onChange={(val) => setJob(`deploy-${cluster}`, val)}
                  env={cluster}
                  chart={chart}
                  images={images}
                  namespace={namespace}
                  releaseName={releaseName}
                />
              </React.Fragment>
            );
          })}
        </CardContent>
      </Card>
    </Container>
  );
};

export default Pipeline;
