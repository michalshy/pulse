import { BrowserRouter, Routes, Route } from "react-router-dom";
import ProjectPage from './views/ProjectPage'
import SessionPage from './views/SessionPage'
import HomePage from './views/HomePage'

export default function App() {
  return (
    <BrowserRouter>
      <Routes>
        <Route path="/" element={<HomePage/>}/>
        <Route path="/project/:id" element={<ProjectPage/>}/>
        <Route path="/session/:id" element={<SessionPage/>}/>
      </Routes>
    </BrowserRouter>
  )
}