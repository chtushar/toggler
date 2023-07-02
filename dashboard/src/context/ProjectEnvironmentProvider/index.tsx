import useProjectEnvironments from '@/hooks/queries/useProjectEnvironments'
import { ApiResponse } from '@/types'
import { Environment } from '@/types/models'
import {
  Dispatch,
  SetStateAction,
  createContext,
  useEffect,
  useState,
} from 'react'

export const ProjectEnvironmentContext = createContext<{
  allEnvironments?: ApiResponse<Array<Environment>>
  currentEnvironment?: Environment
  setCurrentEnvironment: Dispatch<SetStateAction<Environment | undefined>>
} | null>(null)

const ProjectEnvironmentProvider = ({
  children,
}: {
  children: React.ReactNode
}) => {
  const { data: allEnvironments } = useProjectEnvironments()
  const [currentEnvironment, setCurrentEnvironment] = useState<
    Environment | undefined
  >(() => {
    return allEnvironments?.data?.[0]
  })

  useEffect(() => {
    if (!currentEnvironment) {
      setCurrentEnvironment(allEnvironments?.data[0])
    }
  }, [allEnvironments, currentEnvironment])

  return (
    <ProjectEnvironmentContext.Provider
      value={{
        allEnvironments,
        currentEnvironment,
        setCurrentEnvironment,
      }}
    >
      {children}
    </ProjectEnvironmentContext.Provider>
  )
}

export default ProjectEnvironmentProvider
