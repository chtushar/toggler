import { Outlet, useNavigate } from 'react-router-dom'
import Layout from '@/components/common/layout'
import SidebarConfigProvider from '@/context/SidebarConfigProvider'
import useUser from '@/hooks/queries/useUser'
import React from 'react'
import useUserOrganizations from '@/hooks/queries/useUserOrganizations'

const Root = () => {
  const { data, isError } = useUser()
  const { data: userOrgs } = useUserOrganizations()
  const navigate = useNavigate()

  React.useEffect(() => {
    if (data || isError) {
      navigate('/login')
    }

    if (
      typeof userOrgs !== 'undefined' &&
      Array.isArray(userOrgs?.data) &&
      userOrgs.data.length > 0
    ) {
      navigate(`/${userOrgs.data[0].uuid}`, {
        replace: true,
      })
    }
  }, [data, isError, userOrgs, navigate])

  return (
    <SidebarConfigProvider>
      <Layout>
        <Outlet />
      </Layout>
    </SidebarConfigProvider>
  )
}

export default Root
