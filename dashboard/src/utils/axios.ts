import axios from "axios"

const instance = axios.create({
    baseURL: "http://localhost:9091",
    timeout: 5000,
    headers: {
        "Content-Type": "application/json",
    },
})

export default instance