import {
  Card,
  CardContent,
  Container,
  FormControl,
  MenuItem,
  Select,
  TextField,
  SelectChangeEvent,
  InputLabel,
  CardActions,
  FormControlLabel,
  Switch,
  CardHeader,
} from "@mui/material";
import IngressRoutes from "./Routes";
import React from "react";
import { useState } from "react";
import MultiSelect from "./MultiSelect";

const Value = (props: ModelProps<NestedPartial<HelmValues>>) => {
  const [image, setImage] = useState<string>("repo/imageName");
  const [policy, setPolicy] = useState<string>("IfNotPresent");

  const [replicaCount, setReplicaCount] = useState<number>(1);

  const handleChangeImages = (e: React.ChangeEvent<HTMLInputElement>) => {
    setImage(e.target.value);
    props.onChange({
      ...props.value,
      image: {
        ...props.value.image,
        repository: e.target.value,
      },
    });
  };

  const handleChangePolicy = (e: SelectChangeEvent) => {
    setPolicy(e.target.value);
    props.onChange({
      ...props.value,
      image: {
        ...props.value.image,
        pullPolicy: e.target.value,
      },
    });
  };

  const handleChangeReplicaCount = (e: SelectChangeEvent) => {
    setReplicaCount(parseInt(e.target.value));
    props.onChange({
      ...props.value,
      replicaCount: parseInt(e.target.value),
    });
  };

  const [ingEnabled, setIngEnabled] = useState<boolean>(false);
  const [scaleEnabled, setScaleEnabled] = useState<boolean>(false);

  const handleIngressRouteChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    setIngEnabled(e.target.checked);
    // if true then add ingressRoute
    // else remote ingressRoute
    if (e.target.checked) {
      props.onChange({
        ...props.value,
        ingressRoute: {
          enabled: e.target.checked,
          routes: [
            "Host(`test-pregod.rss3.dev`) && PathPrefix(`/v0.4.0`) && Headers(`X-Benchmark-Request`, `1`)",
            "Host(`test-pregod.rss3.dev`) && PathPrefix(`/v0.4.0`)",
          ],
          entrypoint: ["web", "websecure"],
        },
      });
    } else {
      const { ingressRoute, ...rest } = props.value;
      props.onChange(rest);
    }
  };

  const handleScaleChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    setScaleEnabled(e.target.checked);
    if (e.target.checked) {
      props.onChange({
        ...props.value,
        autoscaling: {
          enabled: e.target.checked,
          minReplicas: 1,
          maxReplicas: 10,
          targetCPUUtilizationPercentage: 80,
        },
      });
    } else {
      const { autoscaling, ...rest } = props.value;
      props.onChange(rest);
    }
  };

  return (
    <Container style={{ margin: "20px 0 0 20px" }} maxWidth="sm">
      <Card>
        <CardHeader title="Container" />
        <CardContent>
          <FormControl fullWidth>
            <TextField
              label="Image"
              value={image}
              onChange={handleChangeImages}
            />
          </FormControl>
          <FormControl style={{ marginTop: "20px" }} fullWidth>
            <InputLabel>PullPolicy</InputLabel>
            <Select
              label="PullPolicy"
              value={policy}
              onChange={handleChangePolicy}
            >
              <MenuItem value="Always">Always</MenuItem>
              <MenuItem value="IfNotPresent">IfNotPresent</MenuItem>
              <MenuItem value="Never">Never</MenuItem>
            </Select>
          </FormControl>
          <FormControl style={{ marginTop: "20px" }} fullWidth>
            <InputLabel>ReplicaCount</InputLabel>
            <Select
              label="ReplicaCount"
              value={replicaCount.toString()}
              onChange={handleChangeReplicaCount}
            >
              {[...Array(10)].map((_, i) => (
                <MenuItem key={i} value={(i + 1).toString()}>
                  {i + 1}
                </MenuItem>
              ))}
            </Select>
          </FormControl>
        </CardContent>
      </Card>
      <Card style={{ marginTop: "20px" }}>
        <CardHeader title="Service" />
        <CardContent>
          <FormControl fullWidth>
            <TextField
              label="Service Port"
              value={props.value.service?.port}
              onChange={(e) => {
                props.onChange({
                  ...props.value,
                  service: {
                    ...props.value.service,
                    port: parseInt(e.target.value),
                  },
                });
              }}
            />
          </FormControl>
          <FormControl style={{ marginTop: "20px" }} fullWidth>
            <InputLabel>Serivce Type</InputLabel>
            <Select
              label="Serivce Type"
              value={props.value.service?.type}
              onChange={(e) => {
                props.onChange({
                  ...props.value,
                  service: {
                    ...props.value.service,
                    type: e.target.value,
                  },
                });
              }}
            >
              <MenuItem value="ClusterIP">ClusterIP</MenuItem>
              <MenuItem value="NodePort">NodePort</MenuItem>
              <MenuItem value="LoadBalancer">LoadBalancer</MenuItem>
            </Select>
          </FormControl>
        </CardContent>
      </Card>
      <Card style={{ marginTop: "20px" }}>
        <CardHeader title="IngressRoute" />
        <CardActions>
          <FormControlLabel
            control={
              <Switch value={ingEnabled} onChange={handleIngressRouteChange} />
            }
            label="Enable"
          />
        </CardActions>
        {ingEnabled ? (
          <CardContent>
            <MultiSelect
              value={props.value.ingressRoute?.entrypoint as string[]}
              onChange={(val) => {
                props.onChange({
                  ...props.value,
                  ingressRoute: {
                    ...props.value.ingressRoute,
                    entrypoint: val,
                  },
                });
              }}
              options={["web", "websecure"]}
              label={"Entrypoint"}
            />
            <IngressRoutes
              value={props.value.ingressRoute?.routes as string[]}
              onChange={(val) => {
                props.onChange({
                  ...props.value,
                  ingressRoute: {
                    ...props.value.ingressRoute,
                    routes: val,
                  },
                });
              }}
            />
          </CardContent>
        ) : (
          <CardContent />
        )}
      </Card>
      <Card style={{ marginTop: "20px" }}>
        <CardHeader title="Scale" />
        <CardActions>
          <FormControlLabel
            control={
              <Switch value={scaleEnabled} onChange={handleScaleChange} />
            }
            label="Enable"
          />
        </CardActions>
        {scaleEnabled ? (
          <CardContent>
            <FormControl fullWidth>
              <TextField
                label="MinReplicas"
                value={props.value.autoscaling?.minReplicas}
                onChange={(e) => {
                  props.onChange({
                    ...props.value,
                    autoscaling: {
                      ...props.value.autoscaling,
                      minReplicas: parseInt(e.target.value),
                    },
                  });
                }}
              />
            </FormControl>
            <FormControl style={{ marginTop: "20px" }} fullWidth>
              <TextField
                label="MaxReplicas"
                value={props.value.autoscaling?.maxReplicas}
                onChange={(e) => {
                  props.onChange({
                    ...props.value,
                    autoscaling: {
                      ...props.value.autoscaling,
                      maxReplicas: parseInt(e.target.value),
                    },
                  });
                }}
              />
            </FormControl>
            <FormControl style={{ marginTop: "20px" }} fullWidth>
              <TextField
                label="TargetCPUUtilizationPercentage"
                value={props.value.autoscaling?.targetCPUUtilizationPercentage}
                onChange={(e) => {
                  props.onChange({
                    ...props.value,
                    autoscaling: {
                      ...props.value.autoscaling,
                      targetCPUUtilizationPercentage: parseInt(e.target.value),
                    },
                  });
                }}
              />
            </FormControl>
          </CardContent>
        ) : (
          <CardContent />
        )}
      </Card>
    </Container>
  );
};

export default Value;
