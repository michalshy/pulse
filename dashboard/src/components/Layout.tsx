import { Outlet, NavLink } from "react-router";
import { Database, FileText, Activity, Bell, BellRing, Layers } from "lucide-react";
import { useEffect } from "react";

const navItems = [
  { path: "/projects", label: "Projects", icon: Layers },
  { path: "/logs", label: "Logs", icon: FileText },
  { path: "/metrics", label: "Metrics", icon: Activity },
  { path: "/alert-rules", label: "Alert Rules", icon: Bell },
  { path: "/alert-events", label: "Alert Events", icon: BellRing },
];

export function Layout() {
  useEffect(() => {
    // Ensure dark mode is applied
    document.documentElement.classList.add("dark");
  }, []);

  return (
    <div className="flex h-screen bg-background dark">
      {/* Sidebar */}
      <aside className="w-64 bg-sidebar border-r border-sidebar-border flex flex-col">
        {/* Logo */}
        <div className="p-6 border-b border-sidebar-border">
          <div className="flex items-center gap-3">
            <div className="w-9 h-9 rounded-lg bg-primary/20 flex items-center justify-center">
              <Activity className="w-5 h-5 text-primary" />
            </div>
            <div>
              <h1 className="text-lg font-semibold text-sidebar-foreground">Pulse</h1>
              <p className="text-xs text-muted-foreground">Observability Platform</p>
            </div>
          </div>
        </div>

        {/* Navigation */}
        <nav className="flex-1 p-4 space-y-1">
          {navItems.map((item) => (
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
        </nav>

        {/* Footer */}
        <div className="p-4 border-t border-sidebar-border">
          <p className="text-xs text-muted-foreground">
            v1.0.0 · Pulse
          </p>
        </div>
      </aside>

      {/* Main Content */}
      <main className="flex-1 overflow-auto">
        <Outlet />
      </main>
    </div>
  );
}