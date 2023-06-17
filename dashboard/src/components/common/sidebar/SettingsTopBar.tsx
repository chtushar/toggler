import { ArrowLeft } from 'lucide-react'
import { Button } from '@/components/ui/button'
import { useNavigate } from 'react-router'

const SettingsTopBar = () => {
  const navigate = useNavigate()
  return (
    <Button
      variant="ghost"
      size="sm"
      onClick={() => {
        navigate('/')
      }}
    >
      <ArrowLeft className="mr-2 h-4 w-4" />
      Settings
    </Button>
  )
}

export default SettingsTopBar
