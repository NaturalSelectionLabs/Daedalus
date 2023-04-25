import React, { useState } from "react";
import "./App.css";
import Editor from "@monaco-editor/react";
import Grid from "@mui/material/Grid";
import Pipeline from "./component/Pipeline";
import yaml from "js-yaml";
import AppBar from "./component/AppBar";
import {
  BrowserRouter as Router,
  Routes,
  Route,
  Link,
  Navigate,
} from "react-router-dom";
import Value from "./component/HelmValue";

function App() {
  const [action, setAction] = useState<GithubActionProps>({
    name: "Build and Deploy appName Workflow",
    on: {
      push: {
        branches: ["main"],
        tags: ["v*"],
      },
    },
    jobs: {},
  });

  const [value, setValue] = useState<NestedPartial<HelmValues>>({
    image: {
      repository: "repo/imageName",
    },
    service: {
      type: "ClusterIP",
      port: 80,
    },
  });

  return (
    <Router>
      <div>
        <AppBar></AppBar>
        <div style={{ height: "20px" }} />
        <Routes>
          <Route path="/" element={<Navigate to="/gh-action" />}></Route>
          <Route
            path="/gh-action"
            element={
              <Grid container spacing={2}>
                <Grid item xs={6}>
                  <Pipeline value={action} onChange={setAction} />
                </Grid>
                <Grid item xs={6}>
                  <Editor
                    height="90vh"
                    defaultLanguage="yaml"
                    value={yaml.dump(action, { noCompatMode: true })}
                    options={{
                      readOnly: true,
                    }}
                  />
                </Grid>
              </Grid>
            }
          ></Route>
          <Route path="dockerfile" element={<h1>Dockerfile</h1>}></Route>
          <Route
            path="helm"
            element={
              <Grid container spacing={2}>
                <Grid item xs={6}>
                  <Value value={value} onChange={setValue} />
                </Grid>
                <Grid item xs={6}>
                  <Editor
                    height="90vh"
                    defaultLanguage="yaml"
                    value={yaml.dump(value, { noCompatMode: true })}
                    options={{
                      readOnly: true,
                    }}
                  />
                </Grid>
              </Grid>
            }
          ></Route>
        </Routes>
      </div>
    </Router>
  );
}

export default App;
