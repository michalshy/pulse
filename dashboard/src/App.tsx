import { RouterProvider } from "react-router-dom";
import { Toaster } from "sonner";
import { router } from "./routes";

export default function App() {
  return (
    <>
      <Toaster theme="dark" />
      <RouterProvider router={router}/>
    </>
  )
}