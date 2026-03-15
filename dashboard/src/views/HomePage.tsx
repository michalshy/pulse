import NewProjectButton from '@/components/NewProjectButton'
import ProjectGrid from '@/components/ProjectGrid'
import { useState, useEffect } from 'react'
import type { Project } from '@/models/project'
import { api } from '@/lib/api'

export default function HomePage() {
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
      <NewProjectButton onProjectCreated={fetchProjects} />
      <div className='max-w-7xl mx-auto p-6'>
        <div className='mt-8'>
          <h1 className='text-[#C9D1D9] text-2xl font-semibold'>Projects</h1>
          <p className='text-[#8B949E] mt-2'>Monitor performance across all your projects</p>
        </div>
        
        <div className='mt-4'>
          <input
            type='text'
            placeholder='Search projects...'
            className='w-full bg-[#161B22] text-[#8B949E] p-4 rounded-2xl outline-none'
            />
        </div>
        <ProjectGrid projects={projects}/>
      </div>
    </div>
  )
}