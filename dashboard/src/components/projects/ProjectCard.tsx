import type { Project } from "@/models/project";
import { useNavigate } from "react-router-dom";

export default function ProjectCard({ project }: {project: Project}) {
    const navigate = useNavigate()

    const isActive = project.last_heartbeat
        ? Date.now() - new Date(project.last_heartbeat).getTime() < 120000
        : false

    return (
        <div onClick={() => navigate(`/project/${project.key}/logs`)} className="group bg-card border border-border rounded-xl p-6
            transition-all duration-300 cursor-pointer hover:border-primary/40
            hover:-translate-y-0.5">
            { /* Header */ }
            <div className="flex items-center justify-between mb-1">
                <h2 className="text-lg font-semibold text-card-foreground">{project.name}</h2>
                <span className={`flex items-center gap-1.5 rounded-full px-3 py-1 text-xs
                transition-all duration-300 border
                ${isActive
                    ? 'bg-[var(--pulse-success-bg)] border-[var(--pulse-success)]/25 text-[var(--pulse-success)]'
                    : 'bg-muted border-border text-muted-foreground'
                }`}
                >
                <span className={`w-1.5 h-1.5 rounded-full
                    ${isActive ? 'bg-[var(--pulse-success)] animate-pulse' : 'bg-muted-foreground'}`}
                />
                {isActive ? 'ACTIVE' : 'INACTIVE'}
                </span>
            </div>
            <p className="text-primary text-sm mb-5">{project.key}</p>

            <hr className="border-border mb-4" />

            <p className="text-sm text-muted-foreground leading-relaxed line-clamp-3">
                {project.description}
            </p>

            <div className="flex justify-between text-xs text-muted-foreground mt-4">
                <span>Retention: {project.retention_days} days</span>
            </div>
        </div>
    )
}