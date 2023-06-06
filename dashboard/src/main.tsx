import React from 'react'
import ReactDOM from 'react-dom/client'
import { QueryClientProvider } from '@tanstack/react-query'
import { queryClient } from './utils/queryClient'
import {
  createBrowserRouter,
  redirect,
  RouterProvider,
} from "react-router-dom";
import Root from './scenes/root.tsx';
import RegisterAdmin from './scenes/resgister-admin.tsx';
import './index.css'
import { getHasAdmin } from './hooks/queries/useHasAdmin.ts';

const router = createBrowserRouter([
  {
    path: "/",
    element: <Root />,
    errorElement: <div>Something went wrong</div>,
    loader: async () => {
      const hasAdmin = await getHasAdmin()
      if (!hasAdmin) {
        return redirect("/register-admin")
      }
      return null
    }
  },
  {
    path: "/register-admin",
    element: <RegisterAdmin />,
  }
]);

ReactDOM.createRoot(document.getElementById('root') as HTMLElement).render(
  <React.StrictMode>
    <QueryClientProvider client={queryClient}>
      <RouterProvider router={router} />
    </QueryClientProvider>
  </React.StrictMode>,
)
