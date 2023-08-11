import axios from '@/utils/axios'
import { useQuery } from '@tanstack/react-query'
import useCurrentOrganization from './useCurrentOrganization'
import { queryKey } from '@/constants/queryKey'
import { ApiResponse } from '@/types'
import { User } from '@/types/models'

export const getOrgMembers =
  ({ orgId }: { orgId: string }) =>
  async () => {
    try {
      const response = await axios.get(`/api/v1/get_team_members/${orgId}`)
      return response.data
    } catch (error) {
      throw new Error('There was an error')
    }
  }

const useOrganizationMembers = () => {
  const currentOrg = useCurrentOrganization()
  return useQuery<ApiResponse<Array<User>>>({
    queryKey: queryKey.organizationMembers(currentOrg?.uuid ?? ''),
    enabled: !!currentOrg?.uuid,
    queryFn: getOrgMembers({
      orgId: currentOrg?.id.toString() ?? '',
    }),
  })
}

export default useOrganizationMembers
