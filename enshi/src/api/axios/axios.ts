import axios from "axios";

export const axiosLocalhost = axios.create(
    {
        baseURL: `https://localhost/api/v1/`,
        withCredentials: true,
        headers: {
            
        }
    }
)

axios.defaults.withCredentials = true;