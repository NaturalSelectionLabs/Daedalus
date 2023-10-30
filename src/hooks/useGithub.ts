import { useAxios } from "./useAxios";
import { Endpoints } from "@octokit/types";

const BASE_URL = "";

export const useGithubOrgs = () => {
  type Data = Endpoints["GET /orgs/{org}"]["response"]["data"];

  const { data, fetch: f } = useAxios<Array<Data>>();

  // console.log(result);
  // const { data, error, loading } = result;

  return {
    data,
    fetch: () => f("/api/github/org", "get"),
  };
};

export const useGithubRepos = () => {
  type Data = Endpoints["GET /orgs/{org}/repos"]["response"]["data"];

  const { data, fetch: f } = useAxios<Data>();

  return {
    data,
    fetch: (org: string) => f(`/api/github/repo/${org}`, "get"),
  };
};
