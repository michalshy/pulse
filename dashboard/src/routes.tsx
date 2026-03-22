import { createBrowserRouter } from "react-router-dom";
import { Layout } from "./components/Layout";
import Projects from "./pages/Projects";
import AlertEvents from "./pages/AlertEvents";
import AlertRules from "./pages/AlertRules";
import Metrics from "./pages/Metrics";
import Logs from "./pages/Logs";
import { ProjectsProvider } from "./context/ProjectsContext";

export const router = createBrowserRouter([
  {
    path: "/",
    element: (
        <ProjectsProvider>
            <Layout />
        </ProjectsProvider>
    ),
    children: [
      { index: true, Component: Projects },
      { path: "project/:key/logs", Component: Logs },
      { path: "project/:key/metrics", Component: Metrics },
      { path: "project/:key/alerts", Component: AlertRules },
      { path: "project/:key/events", Component: AlertEvents },
    ]
  }
])