import * as React from "react";

import HandymanIcon from "@mui/icons-material/Handyman";
import GitHubIcon from "@mui/icons-material/GitHub";
// import componets from @mui/material
import {
  AppBar,
  Box,
  Toolbar,
  IconButton,
  Typography,
  Menu,
  Container,
  Avatar,
  Button,
  Tooltip,
  MenuItem,
} from "@mui/material";

import { Link as RouterLink, useNavigate, useLocation } from "react-router-dom";

const routes: Array<{ title: string; path: string }> = [
  { title: "Github Action", path: "/ui/gh-action" },
  // { title: "Dockerfile", path: "/dockerfile" },
  { title: "Helm", path: "/ui/helm" },
];
const settings = ["Profile", "Account", "Dashboard", "Logout"];

function ResponsiveAppBar() {
  const navigate = useNavigate();
  const location = useLocation();

  return (
    <AppBar position="static">
      <Container maxWidth="xl">
        <Toolbar disableGutters>
          <HandymanIcon sx={{ display: { xs: "none", md: "flex" }, mr: 1 }} />
          <Typography
            variant="h6"
            noWrap
            component="a"
            href="/"
            sx={{
              mr: 2,
              display: { xs: "none", md: "flex" },
              fontFamily: "monospace",
              fontWeight: 700,
              letterSpacing: ".3rem",
              color: "inherit",
              textDecoration: "none",
            }}
          >
            Daedalus
          </Typography>

          <Box sx={{ flexGrow: 1, display: { xs: "none", md: "flex" } }}>
            {routes.map((r) => (
              <Button
                key={r.path}
                sx={{ my: 2, color: "white", display: "block" }}
                onClick={() => navigate(r.path)}
              >
                {r.title}
              </Button>
            ))}
          </Box>

          <Box sx={{ flexGrow: 0 }}>
            <Tooltip title="Github">
              <IconButton
                sx={{ p: 0 }}
                href="https://github.com/NaturalSelectionLabs/Daedalus"
                target="_blank"
              >
                <GitHubIcon color="action" />
              </IconButton>
            </Tooltip>
          </Box>
        </Toolbar>
      </Container>
    </AppBar>
  );
}
export default ResponsiveAppBar;
