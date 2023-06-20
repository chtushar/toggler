import { UserCog, PlusIcon } from 'lucide-react'

import DefaultTopbar from '@/components/common/sidebar/DefaultTopBar'
import SettingsTopBar from '@/components/common/sidebar/SettingsTopBar'
import { SidebarConfigType } from '.'

export const defaultSidebarConfig: SidebarConfigType = {
  key: 'DEFAULT',
  topBar: <DefaultTopbar />,
  items: [
    {
      as: 'a',
      path: '/account/preferences',
      label: 'Preferences',
      icon: <UserCog className="mr-2 h-4 w-4" />,
      section: 'account',
    },
  ],
  sections: [
    {
      id: 'organizations',
      label: 'Organizations',
      sectionCTA: {
        as: 'a',
        path: '/organizations/create',
        label: 'Create Organization',
        icon: <PlusIcon className="mr-2 h-4 w-4" />,
      },
    },
    {
      id: 'account',
      label: 'Account',
    },
  ],
}

export const settingsPageSidebarConfig: SidebarConfigType = {
  key: 'SETTINGS',
  topBar: <SettingsTopBar />,
  items: [],
}
