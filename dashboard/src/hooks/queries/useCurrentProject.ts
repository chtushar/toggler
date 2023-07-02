import { useParams } from 'react-router-dom'
import useCurrentOrganization from './useCurrentOrganization'
import useOrgProjects from './useOrgProjects'
import { useMemo } from 'react'
import { Organization } from '@/types/models'

const useCurrentProject = () => {
  const currentOrg = useCurrentOrganization()
  const { data: projects } = useOrgProjects({
    org: currentOrg as Organization,
  })
  const { projectUuid } = useParams()

  const currentProject = useMemo(() => {
    return projects?.data.find(project => project.uuid === projectUuid)
  }, [projectUuid, projects])

  return currentProject
}
export default useCurrentProject
