import { queryKey } from '@/constants/queryKey'
import axios from '@/utils/axios'
import { useQuery } from '@tanstack/react-query'

export const getUser = async () => {
  try {
    const response = await axios.get('/api/get_user')
    return response.data
  } catch (error: any) {
    throw new Error(error)
  }
}

const useUser = () => {
  return useQuery({
    queryKey: queryKey.user(),
    queryFn: getUser,
  })
}

export default useUser
