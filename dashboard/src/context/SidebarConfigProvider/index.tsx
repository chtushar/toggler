import { Dispatch, createContext, useReducer } from 'react'
import { produce } from 'immer'
import { defaultSidebarConfig } from './sidebar-configs'

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
  type: 'DEFAULT' | 'ADD_ORGANIZATIONS'
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
    case 'ADD_ORGANIZATIONS':
      return produce(state, draft => {
        const orgs = draft.sections?.find(
          section => section.id === 'organizations'
        )
        if (orgs) {
          orgs.items = action.data as Array<any>
        }
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
