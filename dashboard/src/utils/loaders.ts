import { queryKey } from '@/constants/queryKey'
import { getUser } from '@/hooks/queries/useUser'
import { getUserOrganizations } from '@/hooks/queries/useUserOrganizations'
import { ApiResponse } from '@/types'
import { Organization, User } from '@/types/models'
import { QueryClient } from '@tanstack/react-query'
import { redirect } from 'react-router-dom'

const fetchInitialData = async (queryClient: QueryClient) => {
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
  const user = queryClient.getQueryData<ApiResponse<User>>(queryKey.user())

  if (typeof user === 'undefined') {
    return redirect('/login')
  }

  if (!queryClient.getQueryData(queryKey.userOrganizations(user.data.uuid))) {
    const userOrgs = await queryClient.fetchQuery<
      ApiResponse<Array<Organization>>
    >({
      queryKey: queryKey.userOrganizations(user.data.uuid),
      queryFn: getUserOrganizations,
    })

    if (typeof userOrgs === 'undefined' || userOrgs.data === null) {
      return redirect('/organizations/new')
    }
  }

  return null
}

export const loginLoader = (queryClient: QueryClient) => async () => {
  await fetchInitialData(queryClient)
  const user = queryClient.getQueryData(queryKey.user())

  if (user) {
    return redirect('/')
  }

  return null
}
