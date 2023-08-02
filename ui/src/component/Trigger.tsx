import Card from "@mui/material/Card";
import CardContent from "@mui/material/CardContent";
import MultiSelect from "./MultiSelect";

const Trigger = (props: ModelProps<TriggerProps>) => {
  const { branches, tags } = props.value.push;

  const setBranches = (branches: string[]) => {
    props.onChange({
      ...props.value,
      push: {
        ...props.value.push,
        branches,
      },
    });
  };

  const setTags = (tags: string[]) => {
    props.onChange({
      ...props.value,
      push: {
        ...props.value.push,
        tags,
      },
    });
  };

  return (
    <Card sx={{ minWidth: 275 }}>
      <CardContent>
        <MultiSelect
          label="Branches"
          value={branches}
          onChange={setBranches}
          options={["main", "dev", "prod"]}
        />
        <MultiSelect
          label="Tags"
          value={tags}
          onChange={setTags}
          options={["v*", "v1.*", "v1.0.*"]}
        />
      </CardContent>
    </Card>
  );
};

export default Trigger;
