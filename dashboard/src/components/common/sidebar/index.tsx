import useSidebarConfig from '@/context/SidebarConfigProvider/useSidebarConfig'
import { Button } from '@/components/ui/button'
import { Link } from 'react-router-dom'
import { CheckIcon } from 'lucide-react'

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
                {section.items?.map((item, index) => {
                  return (
                    <li key={`${item.label}-${index}`} className="w-full">
                      <Button
                        asChild={item.as === 'a'}
                        variant={item.selected ? 'secondary' : 'ghost'}
                        size="sm"
                        className="w-full justify-start"
                        onClick={item?.onClick}
                      >
                        {item.as === 'a' ? (
                          <Link to={item.path}>
                            {item.icon}
                            <span className="flex w-full items-center justify-between">
                              {item.label}
                              {item.selected && (
                                <CheckIcon className="mr-2 h-4 w-4" />
                              )}
                            </span>
                          </Link>
                        ) : (
                          <>
                            {item.icon}
                            <span className="flex w-full items-center justify-between">
                              {item.label}
                              {item.selected && (
                                <CheckIcon className="mr-2 h-4 w-4" />
                              )}
                            </span>
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
                  variant="outline"
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
