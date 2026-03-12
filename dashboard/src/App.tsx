import { useState, useEffect } from "react"

interface Session {
  id: number
  project_id: string
  started_at: string
  ended_at: string | null
}

export default function App() {
  const [sessions, setSessions] = useState<Session[]>([])

  useEffect(() => {
    fetch("http://localhost:8080/sessions/test_game")
      .then(res => res.json())
      .then(data => setSessions(data))
  }, [])

  return (
    <div className="p-8">
      <h1 className="text-2xl font-bold mb-4">Sessions</h1>
      {sessions.map(session => (
        <div key={session.id} className="border p-4 mb-2">
          <p>{session.project_id}</p>
          <p>{session.started_at}</p>
        </div>
      ))}
    </div>
  )
}