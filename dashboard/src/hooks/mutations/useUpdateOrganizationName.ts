import { ApiResponse } from '@/types'
import { Organization } from '@/types/models'
import axios from '@/utils/axios'
import { useMutation, useQueryClient } from '@tanstack/react-query'
import useUserOrganizations from '../queries/useUserOrganizations'
import { produce } from 'immer'
import { queryKey } from '@/constants/queryKey'
import useUser from '../queries/useUser'

interface UpdateOrganizationName {
  name: string
  orgId: number
}

const useUpdateOrganizationName = () => {
  const { data: userOrgs } = useUserOrganizations()
  const { data: user } = useUser()
  const client = useQueryClient()
  return useMutation({
    mutationFn: async (data: UpdateOrganizationName) => {
      const response = await axios.post('/api/v1/update_organization', data)

      return response.data
    },
    onSuccess: async (data: ApiResponse<Organization>) => {
      if (data.success) {
        const updatedOrganizations = produce(userOrgs, draft => {
          if (draft) {
            for (const org of draft.data) {
              if (org.uuid === data.data.uuid) {
                for (const key of Object.keys(org)) {
                  org[key as keyof typeof org] =
                    data.data[key as keyof typeof org]
                }
                break
              }
            }
          }
        })

        await client.setQueryData(
          queryKey.userOrganizations(user?.data.uuid as string),
          updatedOrganizations
        )
      }
    },
  })
}

export default useUpdateOrganizationName
