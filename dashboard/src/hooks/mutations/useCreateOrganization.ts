import { queryKey } from '@/constants/queryKey'
import { ApiResponse } from '@/types'
import { Organization } from '@/types/models'
import axios from '@/utils/axios'
import { useMutation, useQueryClient } from '@tanstack/react-query'
import { useNavigate } from 'react-router-dom'

interface CreateOrganizationData {
  name: string
}

const useCreateOrganization = () => {
  const client = useQueryClient()
  const navigate = useNavigate()

  return useMutation({
    mutationFn: async (data: CreateOrganizationData) => {
      const response = await axios.post('/api/v1/create_organization', data)
      return response.data
    },
    onSuccess: (data: ApiResponse<Organization>) => {
      if (data.success) {
        client.setQueryData(
          queryKey.organization(String(data.data.id)),
          data.data
        )
        navigate(`/organizations/new/${data.data.id}/project`)
      }
    },
  })
}

export default useCreateOrganization
