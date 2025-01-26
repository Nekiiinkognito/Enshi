import axios from "axios";

const environment = import.meta.env.VITE_ENV || 'development';
// const environment = "docker"
const baseURL = environment === "docker" ? "https://localhost/api/v1/" : "http://localhost:9876/";

export const axiosLocalhost = axios.create(
    {
        baseURL: baseURL,
        withCredentials: true,
        headers: {
            
        }
    }
)

axios.defaults.withCredentials = true;