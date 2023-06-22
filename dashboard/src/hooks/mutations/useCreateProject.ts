import { queryKey } from '@/constants/queryKey'
import { ApiResponse } from '@/types'
import { Project } from '@/types/models'
import axios from '@/utils/axios'
import { useMutation, useQueryClient } from '@tanstack/react-query'
import { useNavigate } from 'react-router-dom'
import useUser from '../queries/useUser'

interface CreateProjectData {
  name: string
  orgId: number
}

const useCreateProject = () => {
  const client = useQueryClient()
  const { data: user } = useUser()
  const navigate = useNavigate()

  return useMutation({
    mutationFn: async (data: CreateProjectData) => {
      const response = await axios.post('/api/v1/create_project', data)
      return response.data
    },
    onSuccess: async (data: ApiResponse<Project>) => {
      if (data.success) {
        client.setQueryData(
          queryKey.project(String(data.data.id), String(user?.data.id)),
          data
        )
      }
    },
  })
}

export default useCreateProject
