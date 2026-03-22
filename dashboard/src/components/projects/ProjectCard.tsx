import type { Project } from "@/models/project";
import { useNavigate } from "react-router-dom";

export default function ProjectCard({ project }: {project: Project}) {
    const navigate = useNavigate()

    return (
        <div onClick={() => navigate(`/project/${project.key}/logs`)} className="group max-w-sm bg-[#0d1117] border border-[#1e2733] rounded-xl p-6
            transition-all duration-300 cursor-pointer hover:border-emerald-500/60
            hover:-translate-y-0.5">
            { /* Header */ }
            <div className="flex items-center justify-between mb-1">
                <h2 className="text-lg font-semibold">{project.name}</h2>
                <span className={`flex items-center gap-1.5 rounded-full px-3 py-1 text-xs
                  transition-all duration-300 border
                  ${project.active
                    ? 'bg-emerald-500/10 border-emerald-500/25 text-emerald-400 group-hover:bg-emerald-500/20'
                    : 'bg-gray-500/10 border-gray-500/25 text-gray-400'
                  }`}>
                    <span className={`w-1.5 h-1.5 rounded-full
                                        ${project.active ? 'bg-emerald-400 animate-pulse' : 'bg-gray-500'}`} />
                    {project.active ? 'ACTIVE' : 'INACTIVE'}
                </span>
            </div>
            { /* Subtitle */ }
            <p className="text-sky-500 text-sm mb-5">{project.key}</p>
            { /* Divider */ }
            <hr className="border-[#1e2733] mb-4"/>
            { /* Description */ }
            <p className="text-sm leading-relaxed line-clamp-3">{project.description}</p>
        </div>
    )
}