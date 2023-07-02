import { useContext } from 'react'
import { ProjectEnvironmentContext } from './index'

const useProjectEnvironmentContext = () => {
  const context = useContext(ProjectEnvironmentContext)

  if (!context) {
    throw new Error(
      'Use useProjectEnvironmentContext inside ProjectEnvironmentContext'
    )
  }

  return context
}

export default useProjectEnvironmentContext
