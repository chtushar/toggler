import { queryKey } from '@/constants/queryKey'
import { ApiResponse } from '@/types'
import { Organization, Project } from '@/types/models'
import axios from '@/utils/axios'
import { useQuery } from '@tanstack/react-query'

export const getOrgProjects =
  ({ orgId }: { orgId: string }) =>
  async () => {
    try {
      const response = await axios.get(`/api/v1/get_org_projects/${orgId}`)
      return response.data
    } catch (error) {
      throw new Error('There was an error')
    }
  }

export const useOrgProjects = ({ org }: { org: Organization }) => {
  return useQuery<ApiResponse<Array<Project>>>({
    queryKey: queryKey.projects(org.uuid),
    queryFn: getOrgProjects({ orgId: org.id }),
    enabled: typeof org !== 'undefined',
  })
}
