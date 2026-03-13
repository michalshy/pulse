import { Activity } from "lucide-react"

export default function TopBar() {
  return (
    <nav className="flex items-center justify-between px-6 py-4 bg-[#0D1117] border-b border-[#21262D]">
      <div className="flex items-center gap-2">
        <Activity className="text-[#10B981]" size={20} />
        <span className="text-white font-semibold text-lg">Pulse</span>
      </div>
    </nav>
  )
}