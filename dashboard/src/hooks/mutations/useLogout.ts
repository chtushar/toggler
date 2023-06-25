import { QueryClient, useMutation, useQueryClient } from '@tanstack/react-query'

import axios from '@/utils/axios'

export const logout = (queryClient: QueryClient) => async () => {
  await axios.post('/api/v1/logout')
  await queryClient.removeQueries()

  window.location.reload()
}

const useLogout = () => {
  const queryClient = useQueryClient()

  return useMutation({
    mutationFn: logout(queryClient),
  })
}

export default useLogout
