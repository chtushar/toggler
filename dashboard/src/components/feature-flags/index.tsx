import { Plus } from 'lucide-react'
import { Button } from '../ui/button'
import CreateFeatureFlagDialog from './CreateFeatureFlagDialog'

const FeatureFlags = () => {
  return (
    <div className="w-full flex flex-col max-w-7xl">
      <div className="w-full flex justify-end">
        <CreateFeatureFlagDialog>
          <Button>
            <Plus className="mr-2 h-4 w-4" />
            <span>New Feature Flag</span>
          </Button>
        </CreateFeatureFlagDialog>
      </div>
    </div>
  )
}

export default FeatureFlags
