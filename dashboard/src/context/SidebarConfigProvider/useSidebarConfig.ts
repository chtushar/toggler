import { useContext } from 'react'
import { SidebarConfigContext } from '.'

const useSidebarConfig = () => {
  const context = useContext(SidebarConfigContext)
  if (!context) {
    throw new Error('use useSidebarConfig inside SidebarConfigProvider')
  }

  return context
}

export default useSidebarConfig
