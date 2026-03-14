import { useEffect, useState } from "react"
import { api } from "@/lib/api"
import type { Project } from "@/models/project"

export default function ProjectGrid() {
    const [projects, setProjects] = useState<Project[]>([])

    useEffect(() => {
        api.get('/projects')
            .then((data: Project[]) => setProjects(data))
    }, [])

    return (
        <div className="grid grid-cols-3 gap-4 mt-6">
        </div>
    )
}