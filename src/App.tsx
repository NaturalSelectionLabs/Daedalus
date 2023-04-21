import React, { useState } from "react";
import "./App.css";
import Editor from "@monaco-editor/react";
import Grid from "@mui/material/Grid";
import Pipeline from "./component/Pipeline";
import yaml from "js-yaml";
import AppBar from "./component/AppBar";
import { BrowserRouter as Router, Routes, Route, Link } from "react-router-dom";

function App() {
  const [action, setAction] = useState<GithubActionProps>({
    name: "Example Action",
    on: {
      push: {
        branches: ["main"],
        tags: ["v*"],
      },
    },
    jobs: {},
  });

  return (
    <Router>
      <div>
        <AppBar></AppBar>
        <div style={{ height: "20px" }} />
        <Routes>
          <Route path="/" element={<h1>Home</h1>}></Route>
          <Route
            path="/github-action"
            element={
              <Grid container spacing={2}>
                <Grid item xs={6}>
                  <Pipeline value={action} onChange={setAction} />
                </Grid>
                <Grid item xs={6}>
                  <Editor
                    height="90vh"
                    defaultLanguage="yaml"
                    value={yaml.dump(action)}
                    options={{
                      readOnly: true,
                    }}
                  />
                </Grid>
              </Grid>
            }
          ></Route>
          <Route path="dockerfile" element={<h1>Dockerfile</h1>}></Route>
          <Route path="helm" element={<h1>Helm</h1>}></Route>
        </Routes>
      </div>
    </Router>
  );
}

export default App;
