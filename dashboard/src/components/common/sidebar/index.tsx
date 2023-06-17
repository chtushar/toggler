import useSidebarConfig from '@/context/SidebarConfigProvider/useSidebarConfig'

const Sidebar = () => {
  const { config } = useSidebarConfig()

  return (
    <div className="h-full p-4 md:max-w-[240px] w-full border border-r border-solid border-slate-200">
      {config.topBar}
    </div>
  )
}

export default Sidebar
