import {
  Card,
  CardActions,
  CardContent,
  FormControlLabel,
  Switch,
  Select,
  MenuItem,
  SelectChangeEvent,
  FormControl,
  InputLabel,
} from "@mui/material";

import React from "react";

const label = { inputProps: { "aria-label": "Switch demo" } };

const Deploy = (
  props: ModelProps<JobProps | undefined> & { env: string; namespace: string }
) => {
  const [enable, setEnable] = React.useState(false);

  const value: JobProps = {
    if:
      props.env === "prod"
        ? "startsWith(github.ref, 'refs/tags/v')"
        : undefined,
    uses: "NaturalSelectionLabs/Daedalus/.github/workflows/deploy-tpl.yaml@main",
    needs: props.env === "dev" ? ["build"] : ["build", "deploy-dev"],

    with: {
      images: "imageName",
      cluster: props.env,
      namespace: props.namespace,
    },
    secrets: "inherit",
  };

  const initValue = () => {
    // if props.value is empty, set default value
    if (props.value === undefined || props.value === null) {
      props.onChange(value);
    }
  };

  const handleChange = (event: React.ChangeEvent<HTMLInputElement>) => {
    if (event.target.checked) {
      initValue();
    } else {
      props.onChange(undefined);
    }
    setEnable(event.target.checked);
  };

  const variables = props.value?.with as {
    namespace: string;
  };

  return (
    <Card sx={{ minWidth: 275 }}>
      <CardActions>
        <FormControlLabel
          control={<Switch value={enable} {...label} onChange={handleChange} />}
          label={`Deploy to ${props.env}`}
        />
      </CardActions>
      <CardContent></CardContent>
    </Card>
  );
};

export default Deploy;
