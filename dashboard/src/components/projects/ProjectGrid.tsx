import type { Project } from "@/models/project"
import ProjectCard from "./ProjectCard"

export default function ProjectGrid({ projects } : { projects: Project[] }) {
    return (
        <div className="grid grid-cols-3 gap-4 mt-6">
            {projects.map((project) => (
                <ProjectCard project={project}/>
            ))}
        </div>
    )
}