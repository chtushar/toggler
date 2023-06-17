import { useNavigate } from 'react-router'
import { LogOut, Settings } from 'lucide-react'
import { Avatar, AvatarFallback } from '@/components/ui/avatar'
import {
  DropdownMenu,
  DropdownMenuTrigger,
  DropdownMenuContent,
  DropdownMenuItem,
  DropdownMenuSeparator,
} from '@/components/ui/dropdown-menu'

import useLogout from '@/hooks/mutations/useLogout'
import useUser from '@/hooks/queries/useUser'

const dropdownItems = [
  {
    id: 'settings',
    path: '/settings',
    label: 'Settings',
    icon: <Settings className="mr-2 h-4 w-4" />,
  },
]

const DefaultTopbar = () => {
  const { data } = useUser()
  const { mutate: logout } = useLogout()
  const navigate = useNavigate()

  const handleLogout = (e: Event) => {
    e.preventDefault()
    logout()
  }

  return (
    <div className="py-2 flex justify-end">
      <DropdownMenu>
        <DropdownMenuTrigger
          asChild
          className="outline-none border border-solid border-transparent focus:border-blue-400 rounded-full"
        >
          <button>
            <Avatar>
              {/* <AvatarImage src="https://github.com/chtushar.png"></AvatarImage> */}
              <AvatarFallback>
                {data?.data.name[0].toUpperCase()}
              </AvatarFallback>
            </Avatar>
          </button>
        </DropdownMenuTrigger>
        <DropdownMenuContent className="w-56 pointer-events-auto">
          {/* <DropdownMenuItem
            onSelect={() => {
              navigate('/setting')
            }}
          >
            <Settings className="mr-2 h-4 w-4" />
            Settings
          </DropdownMenuItem> */}
          {dropdownItems.map(item => {
            return (
              <DropdownMenuItem
                key={item.id}
                onSelect={() => {
                  navigate(item.path)
                }}
              >
                {item.icon}
                {item.label}
              </DropdownMenuItem>
            )
          })}
          <DropdownMenuSeparator />
          <DropdownMenuItem
            onSelect={handleLogout}
            className="pointer-events-auto"
            asChild
          >
            <button className="w-full">
              <LogOut className="mr-2 h-4 w-4" />
              Log out
            </button>
          </DropdownMenuItem>
        </DropdownMenuContent>
      </DropdownMenu>
    </div>
  )
}

export default DefaultTopbar
