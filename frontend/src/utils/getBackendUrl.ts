import axios from "axios";

const getBackendUrl = () => {
  return import.meta.env.VITE_BACKEND_URL;
};

export default getBackendUrl;
export const getBackendAxiosInstance = () => {
  return axios.create({
    baseURL: getBackendUrl(),
    headers: {
      "Content-Type": "application/json",
    },
  });
};