import { SidebarConfigType } from '.'
import DefaultTopbar from '@/components/common/sidebar/DefaultTopBar'

export const defaultSidebarConfig: SidebarConfigType = {
  key: 'DEFAULT',
  topBar: <DefaultTopbar />,
  sections: [],
}

export const settingsPageSidebarConfig: SidebarConfigType = {
  key: 'SETTINGS',
  sections: [],
}
