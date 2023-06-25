import { Dispatch, createContext, useReducer } from 'react'
import { produce } from 'immer'
import { defaultSidebarConfig } from './sidebar-configs'
import { PieChart, Braces, Cog } from 'lucide-react'

interface ButtonItem {
  as: 'button'
  label: string
  icon?: React.ReactElement
  onClick: () => void
  path?: never
  section?: string
}

interface AnchorItem {
  as: 'a'
  label: string
  icon?: React.ReactElement
  onClick?: never
  path: string
  section?: string
}

interface SectionItem {
  id: string
  label?: string
  sectionCTA?: ButtonItem | AnchorItem
  items?: Array<ButtonItem | AnchorItem>
}

export interface SidebarConfigType {
  key: string
  topBar?: React.ReactElement
  sections?: Array<SectionItem>
}

export interface SidebarConfigAction {
  type: 'DEFAULT' | 'ORGANIZATION' | 'ORG-PROJECTS'
  data: unknown
}

const SidebarConfigContext = createContext<{
  config: SidebarConfigType
  dispatch: Dispatch<SidebarConfigAction> | null
} | null>(null)

const reducer = (state: SidebarConfigType, action: SidebarConfigAction) => {
  const { orgUuid, projects } = action.data as {
    orgUuid: string
    projects: Array<AnchorItem | ButtonItem>
  }
  switch (action.type) {
    case 'DEFAULT':
      return defaultSidebarConfig
    case 'ORGANIZATION':
      return produce(state, draft => {
        draft.sections = [
          {
            id: 'misc',
            items: [
              {
                as: 'a',
                label: 'Overview',
                path: `/${orgUuid}/overview`,
                icon: <PieChart className="mr-2 h-4 w-4" />,
              },
              {
                as: 'a',
                label: 'Tokens',
                path: `/${orgUuid}/tokens`,
                icon: <Braces className="mr-2 h-4 w-4" />,
              },
              {
                as: 'a',
                label: 'Settings',
                path: `/${orgUuid}/settings`,
                icon: <Cog className="mr-2 h-4 w-4" />,
              },
            ],
          },
          {
            id: 'projects',
            label: 'Projects',
            items: [],
          },
        ]
        return draft
      })
    case 'ORG-PROJECTS':
      return produce(state, draft => {
        const draftProjects = draft.sections?.find(
          section => section.id === 'projects'
        )
        if (typeof draftProjects === 'undefined') {
          draft.sections?.splice(1, 0, {
            id: 'projects',
            label: 'Projects',
            items: projects,
          })
        } else {
          draftProjects.items = projects
        }
      })
    default:
      return defaultSidebarConfig
  }
}

const SidebarConfigProvider = ({ children }: { children: React.ReactNode }) => {
  const [config, dispatch] = useReducer(reducer, defaultSidebarConfig)

  return (
    <SidebarConfigContext.Provider
      value={{
        config,
        dispatch,
      }}
    >
      {children}
    </SidebarConfigContext.Provider>
  )
}

export { SidebarConfigContext }

export default SidebarConfigProvider
