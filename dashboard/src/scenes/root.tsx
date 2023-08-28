import { useEffect } from 'react'
import { Outlet, useLocation, useNavigate } from 'react-router-dom'
import Layout from '@/components/common/layout'
import SidebarConfigProvider from '@/context/SidebarConfigProvider'
import useUserOrganizations from '@/hooks/queries/useUserOrganizations'

const Root = () => {
  const location = useLocation()
  const navigate = useNavigate()
  const { data: userOrgs } = useUserOrganizations()

  useEffect(() => {
    if (location.pathname === '/') {
      navigate(`/${userOrgs?.data[0].uuid}/`)
    }
  }, [location.pathname, navigate, userOrgs?.data])

  return (
    <SidebarConfigProvider>
      <Layout>
        <Outlet />
      </Layout>
    </SidebarConfigProvider>
  )
}

export default Root
