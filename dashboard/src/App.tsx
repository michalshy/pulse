import { BrowserRouter, Routes, Route } from "react-router-dom";

function App() {
  return (
    <BrowserRouter>
      <Routes>
        <Route path="/" element={<ProjectPage/>}/>
        <Route path="/project/:id" element={<SessionPage/>}/>
      </Routes>
    </BrowserRouter>
  )
}