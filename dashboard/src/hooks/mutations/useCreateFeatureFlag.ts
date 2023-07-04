import axios from '@/utils/axios'
import { useMutation } from '@tanstack/react-query'

interface CreateFeatureFlagData {
  name: string
  flag_type: 'boolean'
  project_id: number
}

const useCreateFeatureFlag = () => {
  return useMutation({
    mutationFn: async (data: CreateFeatureFlagData) => {
      const response = await axios.post('/api/v1/create_feature_flag', data)
      return response.data
    },
  })
}

export default useCreateFeatureFlag
