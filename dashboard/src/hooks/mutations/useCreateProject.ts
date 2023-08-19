import { queryKey } from '@/constants/queryKey'
import { ApiResponse } from '@/types'
import { Project } from '@/types/models'
import axios from '@/utils/axios'
import { useMutation, useQueryClient } from '@tanstack/react-query'
import useUser from '../queries/useUser'
import { useNavigate, useParams } from 'react-router-dom'

interface CreateProjectData {
  name: string
  orgId: number
}

const useCreateProject = () => {
  const client = useQueryClient()
  const { data: user } = useUser()
  const { orgUuid } = useParams()
  const navigate = useNavigate()

  return useMutation({
    mutationFn: async (data: CreateProjectData) => {
      const response = await axios.post(
        `/api/v1/create_project/${data.orgId}`,
        {
          name: data.name,
        }
      )
      return response.data
    },
    onSuccess: async (data: ApiResponse<Project>) => {
      if (data.success) {
        client.invalidateQueries({
          queryKey: queryKey.project(
            String(data.data.id),
            String(user?.data.id)
          ),
        })
        if (data?.data?.uuid) {
          navigate(`/${orgUuid}/project/${data.data.uuid}`)
        }
      }
    },
  })
}

export default useCreateProject
