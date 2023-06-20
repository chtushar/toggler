import useSidebarConfig from '@/context/SidebarConfigProvider/useSidebarConfig'
import { Button } from '@/components/ui/button'
import { Link } from 'react-router-dom'

const Sidebar = () => {
  const { config } = useSidebarConfig()

  return (
    <div className="h-full p-4 md:max-w-[240px] w-full border border-r border-solid border-slate-200">
      {config.topBar}
      <div className="flex w-full flex-col">
        {config.sections?.map(section => {
          return (
            <div className="w-full flex flex-col gap-4 py-4">
              {!!section.label && (
                <p className="text-base text-muted-foreground">
                  {section.label}
                </p>
              )}
            </div>
          )
        })}
      </div>
    </div>
  )
}

export default Sidebar
