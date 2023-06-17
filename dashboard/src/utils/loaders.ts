import { queryKey } from '@/constants/queryKey'
import { getHasAdmin } from '@/hooks/queries/useHasAdmin'
import { getUser } from '@/hooks/queries/useUser'
import { QueryClient } from '@tanstack/react-query'
import { redirect } from 'react-router-dom'

const fetchInitialData = async (queryClient: QueryClient) => {
  if (!queryClient.getQueryData(queryKey.hasAdmin())) {
    await queryClient
      .fetchQuery({
        queryKey: queryKey.hasAdmin(),
        queryFn: getHasAdmin,
      })
      .catch(() => {
        return undefined
      })
  }
  if (!queryClient.getQueryData(queryKey.user())) {
    await queryClient
      .fetchQuery({
        queryKey: queryKey.user(),
        queryFn: getUser,
      })
      .catch(() => {
        return undefined
      })
  }
}

export const rootLoader = (queryClient: QueryClient) => async () => {
  await fetchInitialData(queryClient)
  const hasAdmin = queryClient.getQueryData(queryKey.hasAdmin())
  const user = queryClient.getQueryData(queryKey.user())

  if (!hasAdmin) {
    return redirect('/register-admin')
  }

  if (hasAdmin && typeof user === 'undefined') {
    return redirect('/login')
  }

  return null
}

export const loginLoader = (queryClient: QueryClient) => async () => {
  await fetchInitialData(queryClient)
  const hasAdmin = queryClient.getQueryData(queryKey.hasAdmin())
  const user = queryClient.getQueryData(queryKey.user())

  if (!hasAdmin) {
    return redirect('/register-admin')
  }

  if (user) {
    return redirect('/')
  }

  return null
}
