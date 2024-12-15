import axios from "axios";

export const axiosLocalhost = axios.create(
    {
        baseURL: `http://127.0.0.1:9876/`,
        withCredentials: true,
        headers: {
            
        }
    }
)

axios.defaults.withCredentials = true;