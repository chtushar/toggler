import './index.css'
import React from 'react'
import ReactDOM from 'react-dom/client'
import { QueryClientProvider } from '@tanstack/react-query'
import { queryClient } from './utils/queryClient'
import { createBrowserRouter, RouterProvider } from 'react-router-dom'

import Root from './scenes/root.tsx'
import Register from './scenes/register.tsx'
import Login from './scenes/login.tsx'

import { ReactQueryDevtools } from '@tanstack/react-query-devtools'
import { loginLoader, rootLoader } from './utils/loaders.ts'
import New from './scenes/new.tsx'
import CreateOrg from './components/new/CreateOrg.tsx'
import CreateProject from './components/new/CreateProject.tsx'
import Organization from './scenes/organization.tsx'
import OrganizationOverview from './components/organization/index.tsx'
import Tokens from './components/organization/Tokens.tsx'
import Settings from './components/organization/Settings/index.tsx'
import Project from './components/project/index.tsx'

const router = createBrowserRouter([
  {
    path: '/',
    element: <Root />,
    errorElement: <div>Here went wrong</div>,
    loader: rootLoader(queryClient),
    children: [
      {
        path: '/:orgUuid',
        element: <Organization />,
        children: [
          {
            path: '/:orgUuid',
            element: <OrganizationOverview />,
          },
          {
            path: '/:orgUuid/tokens',
            element: <Tokens />,
          },
          {
            path: '/:orgUuid/settings',
            element: <Settings />,
          },
          {
            path: '/:orgUuid/project/:projectUuid',
            element: <Project />,
          },
        ],
      },
    ],
  },
  {
    path: '/organizations',
    element: <New />,
    errorElement: <div>Something went wrong</div>,
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
    element: <Register />,
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
