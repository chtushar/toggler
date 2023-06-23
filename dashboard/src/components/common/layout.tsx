import React from 'react'
import Sidebar from './sidebar'

import useSidebarConfig from '@/context/SidebarConfigProvider/useSidebarConfig'
import useUserOrganizations from '@/hooks/queries/useUserOrganizations'

const Layout = ({ children }: { children: React.ReactNode }) => {
  const { data: userOrgs, isSuccess: isFetchUserOrgsSuccess } =
    useUserOrganizations()
  const { dispatch } = useSidebarConfig()

  React.useEffect(() => {
    if (isFetchUserOrgsSuccess && dispatch) {
      dispatch({
        type: 'ADD_ORGANIZATIONS',
        data: userOrgs.data.map(org => {
          return {
            as: 'a',
            path: '/organizations/' + org.id,
            label: org.name,
          }
        }),
      })
    }
  }, [isFetchUserOrgsSuccess, userOrgs, dispatch])

  return (
    <div className="w-full h-full hidden md:flex">
      <Sidebar />
      <main className="flex-1">{children}</main>
    </div>
  )
}

export default Layout
