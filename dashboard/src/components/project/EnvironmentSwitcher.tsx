import { ChevronsUpDown } from 'lucide-react'
import useProjectEnvironmentContext from '@/context/ProjectEnvironmentProvider/useProjectEnvironmentContext'
import {
  DropdownMenu,
  DropdownMenuContent,
  DropdownMenuGroup,
  DropdownMenuItem,
  DropdownMenuTrigger,
} from '../ui/dropdown-menu'
import { Button } from '../ui/button'

const EnvironmentSwitcher = () => {
  const { currentEnvironment, setCurrentEnvironment, allEnvironments } =
    useProjectEnvironmentContext()

  return (
    <div className="w-48 p-1 border border-dashed border-yellow-600 rounded-md">
      <DropdownMenu>
        <DropdownMenuTrigger asChild>
          <Button
            variant="outline"
            role="combobox"
            className="flex w-full h-fit items-center"
          >
            {currentEnvironment?.name}
            <ChevronsUpDown className="ml-auto h-4 w-4 shrink-0 opacity-50" />
          </Button>
        </DropdownMenuTrigger>
        <DropdownMenuContent className="w-48 p-1">
          <DropdownMenuGroup className="w-full">
            {allEnvironments?.data.map(env => {
              return (
                <DropdownMenuItem
                  className="w-full"
                  key={env.uuid}
                  onSelect={() => setCurrentEnvironment(env)}
                >
                  {env.name}
                </DropdownMenuItem>
              )
            })}
          </DropdownMenuGroup>
        </DropdownMenuContent>
      </DropdownMenu>
    </div>
  )
}

export default EnvironmentSwitcher
