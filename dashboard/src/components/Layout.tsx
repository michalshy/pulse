import { Outlet, NavLink, useParams, useNavigate } from "react-router";
import { FileText, Activity, Bell, BellRing, Layers } from "lucide-react";
import { useEffect } from "react";
import { useProjects } from "@/context/ProjectsContext";

export function Layout() {
  const { projects } = useProjects();
  const { key } = useParams();
  const navigate = useNavigate()

  useEffect(() => {
    document.documentElement.classList.add("dark");
  }, []);

  const projectNavItems = key
    ? [
        { path: `/project/${key}/logs`, label: "Logs", icon: FileText },
        { path: `/project/${key}/metrics`, label: "Metrics", icon: Activity },
        { path: `/project/${key}/alerts`, label: "Alert rules", icon: Bell },
        { path: `/project/${key}/events`, label: "Alert events", icon: BellRing },
      ]
    : [];

  return (
    <div className="flex h-screen bg-background dark">
      <aside className="w-64 bg-sidebar border-r border-sidebar-border flex flex-col">
        <div className="p-6 border-b border-sidebar-border">
          <div className="flex items-center gap-3">
            <div className="w-9 h-9 rounded-lg bg-primary/20 flex items-center justify-center">
              <Activity className="w-5 h-5 text-primary" />
            </div>
            <div>
              <h1 className="text-lg font-semibold text-sidebar-foreground">Pulse</h1>
              <p className="text-xs text-muted-foreground">Observability platform</p>
            </div>
          </div>
        </div>

        <nav className="flex-1 p-4 space-y-1">
          <NavLink
            to="/"
            end
            className={({ isActive }) =>
              `flex items-center gap-3 px-3 py-2.5 rounded-lg transition-colors ${
                isActive && !key
                  ? "bg-sidebar-accent text-sidebar-foreground"
                  : "text-muted-foreground hover:bg-sidebar-accent/50 hover:text-sidebar-foreground"
              }`
            }
          >
            <Layers className="w-5 h-5" />
            <span className="text-sm">Projects</span>
          </NavLink>

          {key && (
            <>
              <div className="pt-4 pb-2 px-3">
                <p className="text-xs text-muted-foreground uppercase tracking-wider">
                  {key}
                </p>
              </div>
              {projectNavItems.map((item) => (
                <NavLink
                  key={item.path}
                  to={item.path}
                  className={({ isActive }) =>
                    `flex items-center gap-3 px-3 py-2.5 rounded-lg transition-colors ${
                      isActive
                        ? "bg-sidebar-accent text-sidebar-foreground"
                        : "text-muted-foreground hover:bg-sidebar-accent/50 hover:text-sidebar-foreground"
                    }`
                  }
                >
                  <item.icon className="w-5 h-5" />
                  <span className="text-sm">{item.label}</span>
                </NavLink>
              ))}
            </>
          )}
        </nav>

        <div className="p-4 border-t border-sidebar-border">
          <p className="text-xs text-muted-foreground">v1.0.0 · Pulse</p>
        </div>
      </aside>

      <main className="flex-1 overflow-auto">
        <Outlet />
      </main>
    </div>
  );
}