import { createBrowserRouter } from "react-router-dom";
import { Layout } from "./components/Layout";
import Projects from "./pages/Projects";

export const router = createBrowserRouter([
    {
        path: "/",
        Component: Layout,
        children: [
            { index: true, Component: Projects }
        ]
    }
])