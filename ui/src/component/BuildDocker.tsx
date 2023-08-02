import FormControlLabel from "@mui/material/FormControlLabel";
import Switch from "@mui/material/Switch";
import Card from "@mui/material/Card";
import CardActions from "@mui/material/CardActions";
import CardContent from "@mui/material/CardContent";
import TextField from "@mui/material/TextField";
import Box from "@mui/material/Box";
import React from "react";
import { FormControl } from "@mui/material";

const label = { inputProps: { "aria-label": "Switch demo" } };

const Job = (
  props: ModelProps<Partial<JobProps> | undefined> & { images: string }
) => {
  const [enable, setEnable] = React.useState(false);

  const value = {
    uses: "NaturalSelectionLabs/Daedalus/.github/workflows/docker-tpl.yaml@main",
    with: {
      images: props.images,
      context: ".",
      dockerfile: "./Dockerfile",
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
    images: string;
    context: string;
    dockerfile: string;
  };

  const setContext = (event: React.ChangeEvent<HTMLInputElement>) => {
    props.onChange({
      ...props.value,
      with: {
        ...props.value?.with,
        context: event.target.value,
      },
    });
  };

  const setDockerfile = (event: React.ChangeEvent<HTMLInputElement>) => {
    props.onChange({
      ...props.value,
      with: {
        ...props.value?.with,
        dockerfile: event.target.value,
      },
    });
  };

  return (
    <Card sx={{ minWidth: 275 }}>
      <CardActions>
        <FormControlLabel
          control={<Switch value={enable} {...label} onChange={handleChange} />}
          label="Build Docker"
        />
      </CardActions>
      <CardContent>
        <FormControl fullWidth>
          <TextField
            label="Context"
            variant="outlined"
            value={(variables || { context: "" }).context}
            onChange={setContext}
            helperText=" "
            disabled={!enable}
          />
          <TextField
            fullWidth
            label="Dockerfile"
            value={(variables || { dockerfile: "" }).dockerfile}
            onChange={setDockerfile}
            helperText=" "
            disabled={!enable}
          />
        </FormControl>
      </CardContent>
    </Card>
  );
};

export default Job;
