import { useEffect, useMemo } from 'react'
import { Outlet, useParams } from 'react-router-dom'
import useSidebarConfig from '@/context/SidebarConfigProvider/useSidebarConfig'

import useUserOrganizations from '@/hooks/queries/useUserOrganizations'
import { useOrgProjects } from '@/hooks/queries/useOrgProjects'

import { Organization as OrganizationModel } from '@/types/models'

const Organization = () => {
  const { dispatch } = useSidebarConfig()
  const { orgUuid, projectUuid } = useParams()
  const { data: userOrgs } = useUserOrganizations()

  const currentOrg = useMemo(() => {
    return userOrgs?.data.find(org => org.uuid === orgUuid)
  }, [orgUuid, userOrgs?.data])

  const { data: orgProjects } = useOrgProjects({
    org: currentOrg as OrganizationModel,
  })

  useEffect(() => {
    if (dispatch) {
      dispatch({
        type: 'ORGANIZATION',
        data: {
          orgUuid,
          projects: orgProjects?.data.map(project => {
            return {
              as: 'a',
              path: `/${orgUuid}/project/${project.uuid}`,
              label: project.name,
              selected: projectUuid === project.uuid,
            }
          }),
        },
      })
    }
  }, [dispatch, orgUuid, orgProjects?.data, projectUuid])

  return (
    <div className="p-4">
      <Outlet />
    </div>
  )
}

export default Organization
