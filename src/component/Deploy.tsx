import Card from "@mui/material/Card";
import CardActions from "@mui/material/CardActions";
import FormControlLabel from "@mui/material/FormControlLabel";
import Switch from "@mui/material/Switch";
import React from "react";

const Deploy = (props: ModelProps<JobProps> & { env: string }) => {
  const [enable, setEnable] = React.useState(true);

  const handleChange = (event: React.ChangeEvent<HTMLInputElement>) => {
    if (event.target.checked) {
      //   initValue();
    } else {
      //   props.onChange(undefined);
    }
    setEnable(event.target.checked);
  };

  return (
    <Card>
      <CardActions>
        <FormControlLabel
          control={<Switch value={enable} onChange={handleChange} />}
          label={`Deploy to ${props.env}`}
        />
      </CardActions>
    </Card>
  );
};

export default Deploy;
