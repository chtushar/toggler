import DefaultTopbar from '@/components/common/sidebar/DefaultTopBar'
import SettingsTopBar from '@/components/common/sidebar/SettingsTopbar'

export const defaultSidebarConfig = {
  key: 'DEFAULT',
  topBar: <DefaultTopbar />,
  items: [],
}

export const settingsPageSidebarConfig = {
  key: 'SETTINGS',
  topBar: <SettingsTopBar />,
  items: [],
}
