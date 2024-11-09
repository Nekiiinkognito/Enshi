import axios from "axios";

export const axiosLocalhost = axios.create(
    {
        baseURL: `http://localhost:9876/`,
        withCredentials: true,
        headers: {
            
        }
    }
)

axios.defaults.withCredentials = true;