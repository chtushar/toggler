import { Dispatch, createContext, useEffect, useReducer } from 'react'
import { Location, useLocation } from 'react-router-dom'
import {
  defaultSidebarConfig,
  settingsPageSidebarConfig,
} from './sidebar-configs'

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
  type: 'DEFAULT' | 'SETTINGS'
  config: SidebarConfigType
}

const SidebarConfigContext = createContext<{
  config: SidebarConfigType
  dispatch: Dispatch<SidebarConfigAction> | null
}>({
  config: defaultSidebarConfig,
  dispatch: null,
})

const reducer = (state: SidebarConfigType, action: SidebarConfigAction) => {
  switch (action.type) {
    case 'DEFAULT':
      return defaultSidebarConfig
    case 'SETTINGS':
      return action.config
    default:
      return defaultSidebarConfig
  }
}

const createInitialState =
  (location: Location) => (initialState: SidebarConfigType) => {
    if (location.pathname === '/settings') {
      return settingsPageSidebarConfig
    }
    return initialState
  }

const SidebarConfigProvider = ({ children }: { children: React.ReactNode }) => {
  const location = useLocation()
  const [config, dispatch] = useReducer(
    reducer,
    defaultSidebarConfig,
    createInitialState(location)
  )

  useEffect(() => {
    if (location.pathname.startsWith('/settings')) {
      dispatch({
        type: 'SETTINGS',
        config: settingsPageSidebarConfig,
      })
    } else {
      dispatch({
        type: 'DEFAULT',
        config: defaultSidebarConfig,
      })
    }
  }, [location])

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
