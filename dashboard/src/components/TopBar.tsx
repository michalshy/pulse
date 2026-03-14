import { Activity, Plus } from "lucide-react"
import { Button } from "./ui/button"
import { useState } from "react"
import NewProjectForm from './NewProjectForm'
import {
  Dialog,
  DialogContent,
  DialogDescription,
  DialogHeader,
  DialogTitle,
} from "@/components/ui/dialog"

export default function TopBar() {
  const [isOpen, setIsOpen] = useState(false)

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
          <DialogContent>
            <DialogHeader>
              <DialogTitle>Add new Project:</DialogTitle>
            </DialogHeader>
            <NewProjectForm/>
          </DialogContent>
        </Dialog>
      </div>
    </nav>
  )
}