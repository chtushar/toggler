import { queryKey } from '@/constants/queryKey'
import { ApiResponse } from '@/types'
import axios from '@/utils/axios'
import { useMutation, useQueryClient } from '@tanstack/react-query'

interface RegisterData {
  name: string
  email: string
  password: string
}

interface RegisterResponseData {
  id: number
  name: string
  email: string
  email_verified: boolean
  role: string
}

const useRegister = () => {
  const client = useQueryClient()

  return useMutation({
    mutationFn: async (data: RegisterData) => {
      const response = await axios.post('/api/v1/add_user', data)
      return response.data
    },
    onSuccess: async (data: ApiResponse<RegisterResponseData>) => {
      if (data.success) {
        await client.setQueryData<RegisterResponseData>(
          queryKey.user(),
          data.data
        )
      }
    },
  })
}

export default useRegister
