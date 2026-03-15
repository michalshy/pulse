import { Activity, Plus } from "lucide-react"
import { Button } from "./ui/button"
import { useState } from "react"
import { toast } from "sonner"
import {
  Dialog,
  DialogContent,
  DialogDescription,
  DialogHeader,
  DialogTitle,
} from "@/components/ui/dialog"
import type { ProjectFormData } from "@/models/project"
import { api } from "@/lib/api"

export default function TopBar() {
  const [isOpen, setIsOpen] = useState(false)

  const onSubmit = async (formData: FormData) => {
    const data: ProjectFormData = {
        key: formData.get("key") as string,
        name: formData.get("name") as string,
        description: formData.get("description") as string,
    }

    const res = await api.post("/projects", data)

    if (res) {
      setIsOpen(false)
      toast.success("Project created successfully!")
    } else {
      toast.error("Failed to create project.")
    }
  }

  return (
    <nav className="bg-[#161B22] border-b border-[#21262D]">
      <div className="flex max-w-7xl justify-between px-6 py-4 mx-auto items-center gap-2">
        <div className="flex items-center gap-3">
          <Activity className="text-[#10B981]" size={42} />
          <span className="text-[#C9D1D9] font-semibold text-lg">Pulse</span>
        </div>
        <Button onClick={() => {setIsOpen(true)}} className="text-[#C9D1D9]">
          <Plus size={16}/>
          New project
        </Button>
        { /* Dialog */ }
        <Dialog open={isOpen} onOpenChange={setIsOpen}>
          <DialogContent className='bg-[]'>
            <DialogHeader>
              <DialogTitle className='text-lg'>Add new Project:</DialogTitle>
            </DialogHeader>
            <DialogDescription className='text-base'>
              Share your innovative thoughts.
            </DialogDescription>
            <form action={onSubmit} className="flex flex-col gap-4">
              <div className="flex flex-col gap-1.5">
                  <label htmlFor="key" className="text-sm">Project key</label>
                  <input 
                      id="key" 
                      name="key" 
                      placeholder="ExG" 
                      className="outline-none rounded-md border px-3 py-2 text-sm" 
                  />
              </div>
              <div className="flex flex-col gap-1.5">
                  <label htmlFor="name" className="text-sm">Project name</label>
                  <input 
                      id="name" 
                      name="name" 
                      placeholder="Example Game" 
                      className="outline-none rounded-md border px-3 py-2 text-sm" 
                  />
              </div>
              <div className="flex flex-col gap-1.5">
                  <label htmlFor="description" className="text-sm">Project description</label>
                  <textarea 
                      id="description" 
                      name="description" 
                      placeholder="What's this project about?" 
                      className="outline-none rounded-md border px-3 py-2 text-sm" 
                  />
              </div>
              <Button type="submit">Add</Button>
            </form>        
          </DialogContent>
        </Dialog>
      </div>
    </nav>
  )
}