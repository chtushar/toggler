import axios from '../utils/axios'
import { redirect } from '@sveltejs/kit';

export const load = async ({
    route,
}) =>  {
    const { data } = await axios.get('/api/has_admin')
    if (data.data === false) {
        if (route.id !== '/register-admin') {
            throw redirect(307, '/register-admin')
        }
    } else if (data.data === true) {
        if (route.id === '/register-admin') {
            throw redirect(307, '/')
        }
    }
}