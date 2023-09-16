import useCurrentOrganization from '@/hooks/queries/useCurrentOrganization'
import useOrgProjects from '@/hooks/queries/useOrgProjects'
import { Organization as OrganizationModel } from '@/types/models'
import { useEffect } from 'react'
import { useNavigate, useParams } from 'react-router-dom'

const Organization = () => {
  const currentOrg = useCurrentOrganization()
  const navigate = useNavigate()
  const { projectUuid } = useParams()
  const { data: orgProjects } = useOrgProjects({
    org: currentOrg as OrganizationModel,
  })

  useEffect(() => {
    if (
      typeof projectUuid === 'undefined' &&
      orgProjects?.success &&
      Array.isArray(orgProjects?.data)
    ) {
      navigate(`/${currentOrg?.uuid}/project/${orgProjects.data?.[0].uuid}`)
    }
  }, [currentOrg, navigate, projectUuid, orgProjects])

  return null
}

export default Organization
