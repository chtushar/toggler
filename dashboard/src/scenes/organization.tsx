import useSidebarConfig from '@/context/SidebarConfigProvider/useSidebarConfig'
import { useEffect } from 'react'

const Organization = () => {
  const { dispatch } = useSidebarConfig()
  useEffect(() => {
    if (dispatch) {
      dispatch({
        type: 'ORGANIZATION',
        data: null,
      })
    }
  }, [dispatch])

  return <div className="p-4"></div>
}

export default Organization
