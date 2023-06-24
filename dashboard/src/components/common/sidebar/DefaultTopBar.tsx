import { useMemo } from 'react'
import { useParams } from 'react-router'
import { LogOut, ChevronsUpDown, Plus, UserCog } from 'lucide-react'
import { Button } from '@/components/ui/button'
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

  const handleLogout = (e: Event) => {
    e.preventDefault()
    logout()
  }

  const currentOrg = useMemo(() => {
    return userOrgs?.data.find(org => org.uuid === orgUuid)
  }, [orgUuid, userOrgs?.data])

  return (
    <div className="flex">
      <DropdownMenu>
        <DropdownMenuTrigger asChild>
          <Button
            variant="outline"
            role="combobox"
            className="flex flex-col w-full h-fit items-start"
          >
            <span className="flex justify-between w-full items-center">
              <span className="truncate line-clamp-1">{currentOrg?.name}</span>
              <ChevronsUpDown className="ml-auto h-4 w-4 shrink-0 opacity-50" />
            </span>
            <span className="text-xs text-muted-foreground font-normal">
              {user?.data.email}
            </span>
          </Button>
        </DropdownMenuTrigger>
        <DropdownMenuContent className="w-56 pointer-events-auto">
          <DropdownMenuGroup>
            <DropdownMenuLabel className="text-muted-foreground font-medium">
              Organizations
            </DropdownMenuLabel>
            {userOrgs?.data.map(org => {
              return (
                <DropdownMenuItem asChild key={`dropdown-${org.uuid}`}>
                  <Link to={`/organizations/${org.uuid}`}>{org.name}</Link>
                </DropdownMenuItem>
              )
            })}
            <DropdownMenuSeparator />
            <DropdownMenuItem asChild>
              <Link to="/organizations/new">
                <Plus className="mr-2 h-4 w-4" />
                <span>Create Organization</span>
              </Link>
            </DropdownMenuItem>
          </DropdownMenuGroup>
          <DropdownMenuSeparator />
          <DropdownMenuGroup>
            <DropdownMenuLabel className="text-muted-foreground font-medium">
              Account
            </DropdownMenuLabel>
            <DropdownMenuItem
              onSelect={handleLogout}
              className="pointer-events-auto w-full"
            >
              <UserCog className="mr-2 h-4 w-4" />
              Preference
            </DropdownMenuItem>
            <DropdownMenuItem
              onSelect={handleLogout}
              className="pointer-events-auto w-full"
            >
              <LogOut className="mr-2 h-4 w-4" />
              Log out
            </DropdownMenuItem>
          </DropdownMenuGroup>
        </DropdownMenuContent>
      </DropdownMenu>
    </div>
  )
}

export default DefaultTopbar
