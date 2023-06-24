import './index.css'
import React from 'react'
import ReactDOM from 'react-dom/client'
import { QueryClientProvider } from '@tanstack/react-query'
import { queryClient } from './utils/queryClient'
import { createBrowserRouter, redirect, RouterProvider } from 'react-router-dom'

import Root from './scenes/root.tsx'
import RegisterAdmin from './scenes/register-admin.tsx'
import Login from './scenes/login.tsx'

import { getHasAdmin } from './hooks/queries/useHasAdmin.ts'
import { ReactQueryDevtools } from '@tanstack/react-query-devtools'
import { loginLoader, rootLoader } from './utils/loaders.ts'
import New from './scenes/new.tsx'
import CreateOrg from './components/new/CreateOrg.tsx'
import CreateProject from './components/new/CreateProject.tsx'
import Organization from './scenes/organization.tsx'

const router = createBrowserRouter([
  {
    path: '/',
    element: <Root />,
    errorElement: <div>Something went wrong</div>,
    loader: rootLoader(queryClient),
    children: [
      {
        path: '/organizations/:orgUuid',
        element: <Organization />,
      },
    ],
  },
  {
    path: '/organizations',
    element: <New />,
    errorElement: <div>Something went wrong</div>,
    loader: rootLoader(queryClient),
    children: [
      {
        path: '/organizations/new',
        element: <CreateOrg />,
      },
      {
        path: '/organizations/new/:orgUuid/project',
        element: <CreateProject />,
      },
    ],
  },
  {
    path: '/login',
    element: <Login />,
    loader: loginLoader(queryClient),
  },
  {
    path: '/register',
    element: <RegisterAdmin />,
  },
])

ReactDOM.createRoot(document.getElementById('root') as HTMLElement).render(
  <React.StrictMode>
    <QueryClientProvider client={queryClient}>
      <RouterProvider router={router} />
      <ReactQueryDevtools initialIsOpen={false} />
    </QueryClientProvider>
  </React.StrictMode>
)
