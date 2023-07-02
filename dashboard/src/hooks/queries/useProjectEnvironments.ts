import { useQuery } from '@tanstack/react-query'
import useCurrentProject from './useCurrentProject'
import { queryKey } from '@/constants/queryKey'
import axios from '@/utils/axios'
import { ApiResponse } from '@/types'
import { Environment } from '@/types/models'

export const getProjectEnvironments = (projectId: string) => async () => {
  const response = await axios.get(
    `/api/v1/get_project_environments/${projectId}`
  )

  return response.data
}

const useProjectEnvironments = () => {
  const currentProject = useCurrentProject()
  return useQuery<ApiResponse<Array<Environment>>>({
    queryKey: queryKey.environments(currentProject?.uuid as string),
    queryFn: getProjectEnvironments(currentProject?.id as string),
    enabled: !!currentProject?.id,
  })
}

export default useProjectEnvironments
