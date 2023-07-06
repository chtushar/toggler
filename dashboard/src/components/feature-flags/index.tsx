import { Plus } from 'lucide-react'
import { Button } from '../ui/button'
import { Table, TableBody, TableCell, TableHead, TableRow } from '../ui/table'
import CreateFeatureFlagDialog from './CreateFeatureFlagDialog'
import { TableHeader } from '../ui/table'

const FeatureFlags = () => {
  return (
    <div className="w-full flex flex-col max-w-7xl">
      <div className="w-full flex flex-col gap-6">
        <div className="w-full flex justify-end">
          <CreateFeatureFlagDialog>
            <Button size="sm">
              <Plus className="mr-2 h-4 w-4" />
              <span>New Feature Flag</span>
            </Button>
          </CreateFeatureFlagDialog>
        </div>
        <Table>
          <TableHeader>
            <TableRow>
              <TableHead>Name</TableHead>
              <TableHead>Enabled</TableHead>
            </TableRow>
            <TableBody>
              <TableRow>
                <TableCell>Dark Mode</TableCell>
              </TableRow>
            </TableBody>
          </TableHeader>
        </Table>
      </div>
    </div>
  )
}

export default FeatureFlags
