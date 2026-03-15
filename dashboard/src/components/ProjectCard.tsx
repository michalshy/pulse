import type { Project } from "@/models/project";
import { useNavigate } from "react-router-dom";

export default function ProjectCard({ project }: {project: Project}) {
    const navigate = useNavigate()

    return (
        <div
            onClick={() => navigate(`/project/${project.id}`)}
            className="cursor-pointer rounded-lg border-[#30363D] bg-[#161B22] p-4 hover:border-[#10B981] transition-colors"
        >
            <div className="flex items-center gap-2 mb-2">
                <span className="text-xs">{project.key}</span>
            </div>
            <span className={`text-xs ${project.active ? "text-[#10B981]" : "text-[#484F58]"}}`}>
                ● {project.active ? "ACTIVE" : "INACTIVE"}
            </span>
            <h3 className="font-medium">{project.name}</h3>
            <p className="text-sm mt-1 line-clamp-2">{project.description}</p>
        </div>
    )
}