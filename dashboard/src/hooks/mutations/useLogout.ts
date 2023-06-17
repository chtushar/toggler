import { useMutation, useQueryClient } from '@tanstack/react-query'

import axios from '@/utils/axios'

const useLogout = () => {
  const queryClient = useQueryClient()

  return useMutation({
    mutationFn: async () => {
      await axios.post('/api/logout')
      queryClient.removeQueries()

      window.location.reload()
    },
  })
}

export default useLogout
