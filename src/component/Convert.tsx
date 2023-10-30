import {
  FormControl,
  Grid,
  InputLabel,
  MenuItem,
  Select,
  SelectChangeEvent,
} from "@mui/material";
import { useEffect, useState } from "react";
import { useGithubOrgs, useGithubRepos } from "../hooks";
import { RepoList } from "../types/github";

const Convert = () => {
  const [org, setOrg] = useState<string>("");

  const [repo, setRepo] = useState<string>("");

  const { data, fetch: fetchOrgs } = useGithubOrgs();

  useEffect(() => {
    fetchOrgs();
  }, [org]);

  const { data: repos, fetch: fetchRepos } = useGithubRepos();

  const handleOrgChange = (o: string) => {
    setOrg(o);
    fetchRepos(o).then(() => {
      console.log(repos);
    });
  };

  return (
    <Grid container spacing={2}>
      <Grid item xs={4}>
        <FormControl style={{ marginTop: "20px" }} fullWidth>
          <InputLabel>Organization</InputLabel>
          <Select
            label="Organization"
            value={org}
            onChange={(e: SelectChangeEvent) => handleOrgChange(e.target.value)}
          >
            {data?.map((item, index) => (
              <MenuItem key={index} value={item.login}>
                {item.login}
              </MenuItem>
            ))}
          </Select>
        </FormControl>
        <FormControl style={{ marginTop: "20px" }} fullWidth>
          <InputLabel>Repo</InputLabel>
          <Select
            label="Repo"
            value={repo}
            onChange={(e: SelectChangeEvent) => setRepo(e.target.value)}
          >
            {repos?.map((item, index) => (
              <MenuItem key={index} value={item.name}>
                {item.name}
              </MenuItem>
            ))}
          </Select>
        </FormControl>
      </Grid>
      <Grid item xs={8}>
        <h2>Org: {org}</h2>
      </Grid>
    </Grid>
  );
};

export default Convert;
