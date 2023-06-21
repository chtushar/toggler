import { useQuery } from '@tanstack/react-query'
import axios from '@/utils/axios'

import { queryKey } from '@/constants/queryKey'

import type { ApiResponse } from '@/types'
import { User } from '@/types/models'

export const getUser = async () => {
  try {
    const response = await axios.get('/api/v1/get_user')
    return response.data
  } catch (error: any) {
    throw new Error(error)
  }
}

const useUser = () => {
  return useQuery<ApiResponse<User>>({
    queryKey: queryKey.user(),
    queryFn: getUser,
  })
}

export default useUser
