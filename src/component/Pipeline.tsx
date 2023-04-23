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
} from "@mui/material";

const clusters = ["dev", "prod"];

const namespaces = ["default", "pregod", "guardian", "tureco"];
const charts = ["web-app", "test"];

const Pipeline = (props: ModelProps<GithubActionProps>) => {
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

  const setJobNamespace = (j: string, n: string) => {
    if (!props.value.jobs[j]) return;
    setJob(j, {
      ...props.value.jobs[j],
      with: {
        ...props.value.jobs[j]?.with,
        namespace: n,
      },
    });
  };

  const setJobChart = (j: string, c: string) => {
    if (!props.value.jobs[j]) return;
    setJob(j, {
      ...props.value.jobs[j],
      with: {
        ...props.value.jobs[j]?.with,
        chart: c,
      },
    });
  };

  const handleChangeNamespace = (event: SelectChangeEvent) => {
    const n = event.target.value as string;
    setNamespace(n);
    // set dev and prod namespace
    setJobNamespace("deploy-dev", n);
    setJobNamespace("deploy-prod", n);
  };

  const handleChangeChart = (event: SelectChangeEvent) => {
    const c = event.target.value as string;
    const chartPath = `./${c}/`;
    setChart(c);
    // set dev and prod chart
    setJobChart("deploy-dev", chartPath);
    setJobChart("deploy-prod", chartPath);
  };

  return (
    <Container style={{ margin: "20px 0 0 20px" }} maxWidth="sm">
      <Trigger value={props.value.on} onChange={setTrigger} />
      <div style={{ marginTop: "20px" }}></div>
      <BuildDocker value={buildJob} onChange={(val) => setJob("build", val)} />
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
                  chart={`./${chart}/`}
                  namespace={namespace}
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
