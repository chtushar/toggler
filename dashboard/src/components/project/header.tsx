import useCurrentProject from '@/hooks/queries/useCurrentProject'
import EnvironmentSwitcher from './EnvironmentSwitcher'

const Header = () => {
  const currentProject = useCurrentProject()
  return (
    <header className="w-full flex justify-between px-4 py-2 border-b border-solid border-slate-200 gap-8 items-center">
      <h1 className="text-sm font-semibold tracking-tight">
        {currentProject?.name}
      </h1>
      <EnvironmentSwitcher />
    </header>
  )
}

export default Header
