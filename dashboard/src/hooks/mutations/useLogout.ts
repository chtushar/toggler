import { useMutation, useQueryClient } from '@tanstack/react-query'
import { useNavigate } from 'react-router-dom'

import axios from '@/utils/axios'

const useLogout = () => {
  const queryClient = useQueryClient()
  const navigate = useNavigate()

  return useMutation({
    mutationFn: async () => {
      await axios.post('/api/logout')
      await queryClient.invalidateQueries()
      navigate('/login')
    },
  })
}

export default useLogout
