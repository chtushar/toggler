import useProjectEnvironmentContext from '@/context/ProjectEnvironmentProvider/useProjectEnvironmentContext'

const Tokens = () => {
  const { currentEnvironment } = useProjectEnvironmentContext()

  return (
    <div>
      {currentEnvironment?.api_keys.map(key => {
        return <div key={key}>{key}</div>
      })}
    </div>
  )
}

export default Tokens
