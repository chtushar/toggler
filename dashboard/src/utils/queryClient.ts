import { QueryClient, QueryCache } from '@tanstack/react-query'

const queryCache = new QueryCache()

export const queryClient = new QueryClient({
  queryCache,
})
