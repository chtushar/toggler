import { queryKey } from '@/constants/queryKey'
import axios from '@/utils/axios'
import { useMutation, useQueryClient } from '@tanstack/react-query'
import useCurrentProject from '../queries/useCurrentProject'
import useProjectEnvironmentContext from '@/context/ProjectEnvironmentProvider/useProjectEnvironmentContext'
import useCurrentOrganization from '../queries/useCurrentOrganization'

const useToggleFeatureFlag = () => {
  const client = useQueryClient()
  const currentOrg = useCurrentOrganization()
  const currentProject = useCurrentProject()
  const { currentEnvironment } = useProjectEnvironmentContext()
  return useMutation({
    mutationFn: async (ffid: string) => {
      const response = await axios.post(
        `/api/v1/toggle_feature_flag/${currentOrg?.id}/${ffid}`
      )
      return response.data
    },
    onSuccess: () => {
      client.invalidateQueries({
        queryKey: queryKey.featureFlags(
          currentProject?.uuid ?? '',
          currentEnvironment?.uuid ?? ''
        ),
      })
    },
  })
}

export default useToggleFeatureFlag
