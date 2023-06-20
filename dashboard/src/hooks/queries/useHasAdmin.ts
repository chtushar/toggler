import { useQuery } from '@tanstack/react-query'
import axios from '../../utils/axios'

import { queryKey } from '../../constants/queryKey'

export const getHasAdmin = async () => {
  try {
    const { data } = await axios.get('/api/v1/has_admin')
    return data.data
  } catch (error: any) {
    throw new Error(error)
  }
}

const useHasAdmin = () => {
  return useQuery({
    queryKey: queryKey.hasAdmin(),
    queryFn: getHasAdmin,
  })
}

export default useHasAdmin
