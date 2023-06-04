import { useQuery } from "@tanstack/react-query"
import axios from "../../utils/axios"

import { queryKey } from "../../constants/queryKey"

const fetcher = async () => {
    const { data } = await axios.get('/api/has_admin')
    return data.data
}

const useHasAdmin = () => {
    return useQuery({
        queryKey: queryKey.hasAdmin(),
        queryFn: fetcher,
    })
}

export default useHasAdmin