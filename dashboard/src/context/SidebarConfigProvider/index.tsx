import { Dispatch, createContext, useReducer } from 'react'
import { produce } from 'immer'
import { defaultSidebarConfig } from './sidebar-configs'
import { PieChart, Users, Braces, Cog } from 'lucide-react'

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
  type: 'DEFAULT' | 'ORGANIZATION'
  data: unknown
}

const SidebarConfigContext = createContext<{
  config: SidebarConfigType
  dispatch: Dispatch<SidebarConfigAction> | null
} | null>(null)

const reducer = (state: SidebarConfigType, action: SidebarConfigAction) => {
  switch (action.type) {
    case 'DEFAULT':
      return defaultSidebarConfig
    case 'ORGANIZATION':
      return produce(defaultSidebarConfig, draft => {
        draft.sections = [
          {
            id: 'misc',
            items: [
              {
                as: 'a',
                label: 'Overview',
                path: `/organizations`,
                icon: <PieChart className="mr-2 h-4 w-4" />,
              },
              {
                as: 'a',
                label: 'Tokens',
                path: `/organizations`,
                icon: <Braces className="mr-2 h-4 w-4" />,
              },
              {
                as: 'a',
                label: 'Settings',
                path: `/organizations`,
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
