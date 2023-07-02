import { useMemo } from 'react'
import { useParams } from 'react-router-dom'
import useUserOrganizations from './useUserOrganizations'

const useCurrentOrganization = () => {
  const { orgUuid } = useParams()
  const { data: userOrgs } = useUserOrganizations()

  const currentOrg = useMemo(() => {
    return userOrgs?.data.find(org => org.uuid === orgUuid)
  }, [orgUuid, userOrgs?.data])

  return currentOrg
}

export default useCurrentOrganization
