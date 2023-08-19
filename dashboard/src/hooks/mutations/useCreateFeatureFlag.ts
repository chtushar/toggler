import { queryKey } from '@/constants/queryKey'
import axios from '@/utils/axios'
import { useMutation, useQueryClient } from '@tanstack/react-query'
import useCurrentProject from '../queries/useCurrentProject'
import useProjectEnvironmentContext from '@/context/ProjectEnvironmentProvider/useProjectEnvironmentContext'

interface CreateFeatureFlagData {
  name: string
  flag_type: 'boolean'
  project_id: number
}

const useCreateFeatureFlag = () => {
  const client = useQueryClient()
  const currenProject = useCurrentProject()
  const { currentEnvironment } = useProjectEnvironmentContext()

  return useMutation({
    mutationFn: async (data: CreateFeatureFlagData) => {
      const response = await axios.post('/api/v1/create_feature_flag', data)
      return response.data
    },
    onSuccess: () => {
      client.invalidateQueries({
        queryKey: queryKey.featureFlags(
          currenProject?.uuid ?? '',
          currentEnvironment?.uuid ?? ''
        ),
      })
    },
  })
}

export default useCreateFeatureFlag
