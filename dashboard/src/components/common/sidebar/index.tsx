import useSidebarConfig from '@/context/SidebarConfigProvider/useSidebarConfig'
import { Button } from '@/components/ui/button'
import { Link } from 'react-router-dom'

const Sidebar = () => {
  const { config } = useSidebarConfig()

  return (
    <div className="h-full p-4 md:max-w-[240px] w-full border-r border-solid border-slate-200">
      {config.topBar}
      <div className="flex w-full flex-col">
        {config.sections?.map(section => {
          return (
            <div
              key={section.id}
              className="w-full flex flex-col gap-4 py-4 border-b border-solid border-muted-background"
            >
              {!!section.label && (
                <p className="text-sm text-muted-foreground">{section.label}</p>
              )}
              <ul className="w-full">
                {section.items?.map(item => {
                  return (
                    <li key={item.label} className="w-full">
                      <Button
                        asChild={item.as === 'a'}
                        variant="ghost"
                        size="sm"
                        className="w-full justify-start"
                        onClick={item?.onClick}
                      >
                        {item.as === 'a' ? (
                          <Link to={item.path}>
                            {item.icon}
                            {item.label}
                          </Link>
                        ) : (
                          <>
                            {item.icon}
                            {item.label}
                          </>
                        )}
                      </Button>
                    </li>
                  )
                })}
              </ul>
              {section.sectionCTA && (
                <Button
                  asChild={section.sectionCTA.as === 'a'}
                  variant="secondary"
                  size="sm"
                  className="w-full"
                  onClick={section.sectionCTA?.onClick}
                >
                  {section.sectionCTA.as === 'a' ? (
                    <Link to={section.sectionCTA.path}>
                      {section.sectionCTA.icon}
                      {section.sectionCTA.label}
                    </Link>
                  ) : (
                    <>
                      {section.sectionCTA.icon}
                      {section.sectionCTA.label}
                    </>
                  )}
                </Button>
              )}
            </div>
          )
        })}
      </div>
    </div>
  )
}

export default Sidebar
