import useSidebarConfig from '@/context/SidebarConfigProvider/useSidebarConfig'
import { Button } from '@/components/ui/button'
import { Link } from 'react-router-dom'

const Sidebar = () => {
  const { config } = useSidebarConfig()

  return (
    <div className="h-full p-4 md:max-w-[240px] w-full border border-r border-solid border-slate-200">
      {config.topBar}
      <ul className="w-full mt-16">
        {config.items.map(item => {
          return (
            <li key={item.label} className="w-full">
              <Button
                asChild={item.as === 'a'}
                className="w-full justify-start"
                variant="ghost"
              >
                {item.as === 'a' ? (
                  <Link to={item.path}>
                    {item?.icon}
                    {item.label}
                  </Link>
                ) : (
                  <>
                    {item?.icon}
                    {item.label}
                  </>
                )}
              </Button>
            </li>
          )
        })}
      </ul>
    </div>
  )
}

export default Sidebar
