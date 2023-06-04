import a from 'axios'

const axios = a.create({
    baseURL: 'http://localhost:9091',
    timeout: 3000,
    headers: {
        'Content-Type': 'application/json',
    },
})

export default axios