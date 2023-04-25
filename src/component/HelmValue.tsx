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
} from "@mui/material";
import IngressRoutes from "./Routes";
import React from "react";
import { useState } from "react";

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
        },
      });
    } else {
      const { ingressRoute, ...rest } = props.value;
      props.onChange(rest);
    }
  };

  return (
    <Container style={{ margin: "20px 0 0 20px" }} maxWidth="sm">
      <Card>
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
      <Card title="IngressRoute">
        <CardActions>
          <FormControlLabel
            control={
              <Switch value={ingEnabled} onChange={handleIngressRouteChange} />
            }
            label="Enable IngressRoute"
          />
        </CardActions>
        {ingEnabled ? (
          <CardContent>
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
    </Container>
  );
};

export default Value;
