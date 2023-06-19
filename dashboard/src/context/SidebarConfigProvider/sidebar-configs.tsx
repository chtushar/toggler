import { UserCog, Users2 } from 'lucide-react'

import DefaultTopbar from '@/components/common/sidebar/DefaultTopBar'
import SettingsTopBar from '@/components/common/sidebar/SettingsTopBar'
import { SidebarConfigType } from '.'

export const defaultSidebarConfig = {
  key: 'DEFAULT',
  topBar: <DefaultTopbar />,
  items: [],
}

export const settingsPageSidebarConfig: SidebarConfigType = {
  key: 'SETTINGS',
  topBar: <SettingsTopBar />,
  items: [
    {
      as: 'a',
      icon: <UserCog className="mr-2 h-4 w-4" />,
      path: '/settings/account',
      label: 'Account',
    },
    {
      as: 'a',
      icon: <Users2 className="mr-2 h-4 w-4" />,
      path: '/settings/members',
      label: 'Members',
    },
  ],
}
