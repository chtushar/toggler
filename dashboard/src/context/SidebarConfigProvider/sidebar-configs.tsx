import { UserCog, PlusIcon, LogOut } from 'lucide-react'

import DefaultTopbar from '@/components/common/sidebar/DefaultTopBar'
import SettingsTopBar from '@/components/common/sidebar/SettingsTopBar'

import { logout } from '@/hooks/mutations/useLogout'
import { queryClient } from '@/utils/queryClient'

import { SidebarConfigType } from '.'

export const defaultSidebarConfig: SidebarConfigType = {
  key: 'DEFAULT',
  sections: [
    {
      id: 'organizations',
      label: 'Organizations',
      sectionCTA: {
        as: 'a',
        path: '/organizations/new',
        label: 'Create Organization',
        icon: <PlusIcon className="mr-2 h-4 w-4" />,
      },
    },
    {
      id: 'account',
      label: 'Account',
      items: [
        {
          as: 'a',
          path: '/account/preferences',
          label: 'Preferences',
          icon: <UserCog className="mr-2 h-4 w-4" />,
          section: 'account',
        },
        {
          as: 'button',
          onClick: logout(queryClient),
          label: 'Log Out',
          icon: <LogOut className="mr-2 h-4 w-4" />,
        },
      ],
    },
  ],
}

export const settingsPageSidebarConfig: SidebarConfigType = {
  key: 'SETTINGS',
  sections: [],
}
