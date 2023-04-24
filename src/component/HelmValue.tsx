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
} from "@mui/material";
import { useState } from "react";

const Value = (props: ModelProps<NestedPartial<HelmValues>>) => {
  const [image, setImage] = useState<string>("repo/imageName");
  const [policy, setPolicy] = useState<string>("IfNotPresent");

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
            <div style={{ marginTop: "20px" }}></div>
          </FormControl>
          <FormControl fullWidth>
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
        </CardContent>
      </Card>
    </Container>
  );
};

export default Value;
