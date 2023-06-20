import axios from '@/utils/axios'
import { queryKey } from '@/constants/queryKey'
import { useMutation, useQueryClient } from '@tanstack/react-query'
import { useNavigate } from 'react-router-dom'

export interface LoginData {
  email: string
  password: string
}

const useLogin = () => {
  const client = useQueryClient()
  const navigate = useNavigate()

  return useMutation({
    mutationFn: async (data: LoginData) => {
      const response = await axios.post('/api/v1/login', data)
      return response.data
    },
    onSuccess: async data => {
      if (data.success) {
        await client.setQueryData(queryKey.user(), data)
        navigate('/')
      }
    },
  })
}

export default useLogin
