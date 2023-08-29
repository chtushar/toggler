import { useQuery } from '@tanstack/react-query'
import useUser from './useUser'
import { queryKey } from '@/constants/queryKey'
import axios from '@/utils/axios'
import { Organization } from '@/types/models'
import { ApiResponse } from '@/types'

export const getUserOrganizations = async () => {
  try {
    const response = await axios.get('/api/v1/get_user_organizations')
    return response.data
  } catch (error: any) {
    throw new Error(error)
  }
}

const useUserOrganizations = () => {
  const { data, isLoading } = useUser()
  return useQuery<ApiResponse<Array<Organization>>>({
    queryKey: queryKey.userOrganizations(data?.data?.uuid as string),
    queryFn: getUserOrganizations,
    enabled: !isLoading,
  })
}

export default useUserOrganizations
