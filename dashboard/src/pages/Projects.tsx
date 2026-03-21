import ProjectButton from '@/components/projects/ProjectButton'
import ProjectGrid from '@/components/projects/ProjectGrid'
import { useState, useEffect } from 'react'
import type { Project } from '@/models/project'
import { api } from '@/lib/api'

export default function Projects() {
  const [projects, setProjects] = useState<Project[]>([])

  const fetchProjects = () => {
    api.get('/projects')
      .then((data: Project[]) => setProjects(data))
  }

  useEffect(() => {
    fetchProjects()
  }, [])

  return (
    <div className="w-full bg-[#0D1117] min-h-screen">
      <div className='max-w-7xl mx-auto p-6'>
        <div className='flex justify-between items-center mt-8'>
          <div>
            <h1 className='text-[#C9D1D9] text-2xl font-semibold'>Projects</h1>
            <p className='text-[#8B949E] mt-2'>Observe projects, check metrics and traces</p>
          </div>
          <div>
            <ProjectButton onProjectCreated={fetchProjects} />
          </div>
        </div>
        <ProjectGrid projects={projects}/>
      </div>
    </div>
  )
}