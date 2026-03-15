import { BrowserRouter, Routes, Route } from "react-router-dom";
import ProjectPage from './views/ProjectPage'
import SessionPage from './views/SessionPage'
import HomePage from './views/HomePage'
import { Toaster } from "sonner";

export default function App() {
  return (
    <BrowserRouter>
      <Toaster theme="dark" />
      <Routes>
        <Route path="/" element={<HomePage/>}/>
        <Route path="/project/:id" element={<ProjectPage/>}/>
        <Route path="/session/:id" element={<SessionPage/>}/>
      </Routes>
    </BrowserRouter>
  )
}