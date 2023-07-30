import { useQuery } from '@tanstack/react-query'
import useCurrentProject from './useCurrentProject'
import useProjectEnvironmentContext from '@/context/ProjectEnvironmentProvider/useProjectEnvironmentContext'
import { queryKey } from '@/constants/queryKey'
import axios from '@/utils/axios'
import { ApiResponse } from '@/types'
import { FeatureFlag } from '@/types/models'

export const getProjectFeatureFlags =
  (projectId: string, environmentId: string) => async () => {
    const response = await axios.get(
      `/api/v1/get_project_feature_flags/${projectId}/${environmentId}`
    )

    return response.data
  }

const useProjectFeatureFlags = () => {
  const currentProject = useCurrentProject()
  const { currentEnvironment } = useProjectEnvironmentContext()

  return useQuery<ApiResponse<Array<FeatureFlag>>>({
    queryKey: queryKey.featureFlags(
      currentProject?.uuid ?? '',
      currentEnvironment?.uuid ?? ''
    ),
    queryFn: getProjectFeatureFlags(
      currentProject?.id ?? '',
      currentEnvironment?.id ?? ''
    ),
    enabled: !!currentProject?.id && !!currentEnvironment?.id,
  })
}

export default useProjectFeatureFlags
