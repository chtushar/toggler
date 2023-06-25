import { useEffect } from 'react'
import useSidebarConfig from '@/context/SidebarConfigProvider/useSidebarConfig'
import { useParams } from 'react-router-dom'

const Organization = () => {
  const { dispatch } = useSidebarConfig()
  const { orgUuid } = useParams()

  useEffect(() => {
    if (dispatch) {
      dispatch({
        type: 'ORGANIZATION',
        data: {
          orgUuid,
        },
      })
    }
  }, [dispatch, orgUuid])

  return <div className="p-4"></div>
}

export default Organization
