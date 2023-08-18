import axios from '@/utils/axios'
import { useMutation, useQueryClient } from '@tanstack/react-query'
import useCurrentOrganization from '../queries/useCurrentOrganization'
import { queryKey } from '@/constants/queryKey'

export interface AddOrganizationMemberData {
  email: string
}

export const addOrganizationMember =
  (orgId: string) => async (data: AddOrganizationMemberData) => {
    const response = await axios.post(`/api/v1/add_team_member/${orgId}`, data)
    return response.data
  }

const useAddOrganizationMember = () => {
  const currentOrg = useCurrentOrganization()
  const client = useQueryClient()
  return useMutation({
    mutationFn: addOrganizationMember(currentOrg?.id ?? ''),
    onSuccess: () => {
      client.invalidateQueries({
        queryKey: queryKey.organizationMembers(currentOrg?.uuid ?? ''),
      })
    },
  })
}

export default useAddOrganizationMember
