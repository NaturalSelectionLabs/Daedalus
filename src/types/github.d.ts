import { Endpoints } from "@octokit/types";
type RepoList = Endpoints["GET /orgs/{org}/repos"]["response"]["data"];
