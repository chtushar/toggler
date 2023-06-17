import { Outlet, redirect } from 'react-router-dom'
import Layout from '@/components/common/layout'
import SidebarConfigProvider from '@/context/SidebarConfigProvider'
import useUser from '@/hooks/queries/useUser'
import React from 'react'

const Root = () => {
  const { data, isError } = useUser()

  React.useEffect(() => {
    if (data || isError) {
      redirect('/login')
    }
  }, [data, isError])

  return (
    <SidebarConfigProvider>
      <Layout>
        <Outlet />
      </Layout>
    </SidebarConfigProvider>
  )
}

export default Root
