import { useMemo } from 'react'
import { useNavigate, useParams, useLocation } from 'react-router'
import { LogOut, Settings } from 'lucide-react'
import { Avatar, AvatarFallback } from '@/components/ui/avatar'
import {
  DropdownMenu,
  DropdownMenuTrigger,
  DropdownMenuContent,
  DropdownMenuItem,
  DropdownMenuSeparator,
  DropdownMenuGroup,
  DropdownMenuLabel,
} from '@/components/ui/dropdown-menu'

import useLogout from '@/hooks/mutations/useLogout'
import useUser from '@/hooks/queries/useUser'
import useUserOrganizations from '@/hooks/queries/useUserOrganizations'
import { Link } from 'react-router-dom'

const DefaultTopbar = () => {
  const { data: user } = useUser()
  const { mutate: logout } = useLogout()
  const { data: userOrgs } = useUserOrganizations()
  const { orgUuid } = useParams()
  const navigate = useNavigate()

  const handleLogout = (e: Event) => {
    e.preventDefault()
    logout()
  }

  const currentOrg = useMemo(() => {
    return userOrgs?.data.find(org => org.uuid === orgUuid)
  }, [orgUuid, userOrgs?.data])

  return (
    <div className="py-2 flex justify-end">
      <DropdownMenu>
        <DropdownMenuTrigger
          asChild
          className="outline-none border border-solid border-transparent focus:border-blue-400 rounded-full"
        >
          <button className="flex flex-col">
            <span>{currentOrg?.name}</span>
            <span>{user?.data.email}</span>
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
          <DropdownMenuGroup>
            <DropdownMenuLabel>{user?.data.email}</DropdownMenuLabel>
            {userOrgs?.data.map(org => {
              return (
                <DropdownMenuItem asChild key={`dropdown-${org.uuid}`}>
                  <Link to={`/organizations/${org.uuid}`}>{org.name}</Link>
                </DropdownMenuItem>
              )
            })}
          </DropdownMenuGroup>
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
