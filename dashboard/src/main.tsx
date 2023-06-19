import './index.css'
import React from 'react'
import ReactDOM from 'react-dom/client'
import { QueryClientProvider } from '@tanstack/react-query'
import { queryClient } from './utils/queryClient'
import { createBrowserRouter, redirect, RouterProvider } from 'react-router-dom'

import Root from './scenes/root.tsx'
import RegisterAdmin from './scenes/resgister-admin.tsx'
import Login from './scenes/login.tsx'
import Settings from './scenes/settings.tsx'

import Members from './components/settings/members.tsx'

import { getHasAdmin } from './hooks/queries/useHasAdmin.ts'
import { ReactQueryDevtools } from '@tanstack/react-query-devtools'
import { loginLoader, rootLoader } from './utils/loaders.ts'

const router = createBrowserRouter([
  {
    path: '/',
    element: <Root />,
    errorElement: <div>Something went wrong</div>,
    loader: rootLoader(queryClient),
    children: [
      {
        path: '/settings',
        element: <Settings />,
        children: [
          {
            path: '/settings/account',
            element: <Members />,
          },
          {
            path: '/settings/members',
            element: <Members />,
          },
        ],
      },
    ],
  },
  {
    path: '/login',
    element: <Login />,
    loader: loginLoader(queryClient),
  },
  {
    path: '/register-admin',
    element: <RegisterAdmin />,
    loader: async () => {
      const hasAdmin = await getHasAdmin()
      if (hasAdmin) {
        return redirect('/')
      }
      return null
    },
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
