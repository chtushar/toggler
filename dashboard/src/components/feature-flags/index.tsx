import { useState } from 'react'
import { Plus } from 'lucide-react'
import { Button } from '../ui/button'
import { Table, TableBody, TableCell, TableHead, TableRow } from '../ui/table'
import CreateFeatureFlagDialog from './CreateFeatureFlagDialog'
import { TableHeader } from '../ui/table'
import useProjectFeatureFlags from '@/hooks/queries/useProjectFeatureFlags'
import useToggleFeatureFlag from '@/hooks/mutations/useToggleFeatureFlag'
import { Switch } from '../ui/switch'

const FeatureFlags = () => {
  const [openModal, setOpenModal] = useState(false)
  const { data, isLoading } = useProjectFeatureFlags()
  const { mutate: toggleFeatureFlag } = useToggleFeatureFlag()

  return (
    <div className="w-full flex flex-col max-w-7xl">
      <div className="w-full flex flex-col gap-6">
        <div className="w-full flex justify-end">
          <CreateFeatureFlagDialog open={openModal} setOpen={setOpenModal}>
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
          </TableHeader>
          <TableBody>
            {!isLoading &&
              data?.data?.map(ff => {
                return (
                  <TableRow key={ff.uuid}>
                    <TableCell>{ff.name}</TableCell>
                    <TableCell>
                      <Switch
                        checked={ff.enabled}
                        onClick={() => {
                          toggleFeatureFlag(ff.id)
                        }}
                      />
                    </TableCell>
                  </TableRow>
                )
              })}
          </TableBody>
        </Table>
      </div>
    </div>
  )
}

export default FeatureFlags
