import { useEffect } from 'react'
import { Outlet, useNavigate, useParams } from 'react-router-dom'
import useSidebarConfig from '@/context/SidebarConfigProvider/useSidebarConfig'

import useOrgProjects from '@/hooks/queries/useOrgProjects'

import { Organization as OrganizationModel } from '@/types/models'
import useCurrentOrganization from '@/hooks/queries/useCurrentOrganization'

const Organization = () => {
  const { dispatch } = useSidebarConfig()
  const { projectUuid } = useParams()
  const currentOrg = useCurrentOrganization()
  const { data: orgProjects } = useOrgProjects({
    org: currentOrg as OrganizationModel,
  })
  const navigate = useNavigate()

  useEffect(() => {
    if (dispatch) {
      dispatch({
        type: 'ORGANIZATION',
        data: {
          orgUuid: currentOrg?.uuid,
          projects: orgProjects?.data.map(project => {
            return {
              as: 'a',
              path: `/${currentOrg?.uuid}/project/${project.uuid}`,
              label: project.name,
              selected: projectUuid === project.uuid,
            }
          }),
        },
      })
    }
  }, [dispatch, currentOrg?.uuid, orgProjects?.data, projectUuid])

  if (typeof projectUuid === 'undefined' && orgProjects?.success) {
    navigate(`/${currentOrg?.uuid}/project/${orgProjects.data[0].uuid}`)
    return <></>
  }

  return (
    <div>
      <Outlet />
    </div>
  )
}

export default Organization
