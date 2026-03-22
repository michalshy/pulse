import { createContext, useContext, useState, useEffect } from "react";
import type { Project } from "@/models/project";
import { api } from "@/lib/api";

type ProjectsContextType = {
    projects: Project[]
    fetchProjects: () => void
}

const ProjectsContext = createContext<ProjectsContextType>({
    projects: [],
    fetchProjects: () => {}
})

export function ProjectsProvider({ children }: { children: React.ReactNode }) {
    const [projects, setProjects] = useState<Project[]>([])

    const fetchProjects = () => {
        api.get("/projects").then((data: Project[]) => setProjects(data))
    }

    useEffect(() => {
        fetchProjects()
    }, [])

    return (
        <ProjectsContext.Provider value={{ projects, fetchProjects }}>
            {children}
        </ProjectsContext.Provider>
    )
}

export const useProjects = () => useContext(ProjectsContext)