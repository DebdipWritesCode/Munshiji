import axios from "axios";
import { store } from "@/store/store";
import { setAccessToken, clearAccessToken } from "@/slices/authSlice";
import { getBackendUrl } from "@/utils/getBackendUrl";

const api = axios.create({
  baseURL: getBackendUrl(),
  withCredentials: true,
});

api.interceptors.request.use((config) => {
  const token = store.getState().auth.accessToken;
  if (token && config.headers) {
    config.headers.Authorization = `Bearer ${token}`;
  }
  return config;
})

api.interceptors.response.use(
  (response) => response,
  async (error) => {
    const originalRequest = error.config;

    if (
      error.response?.status === 401 &&
      !originalRequest._retry &&
      !originalRequest.url.includes("/refresh_access_token")
    ) {
      originalRequest._retry = true;
      try {
        const res = await axios.post(
          `${getBackendUrl()}/refresh_access_token`,
          {},
          {
            withCredentials: true,
          }
        );
        const newAccessToken = res.data.jwt_token;
        store.dispatch(setAccessToken(newAccessToken));
        originalRequest.headers.Authorization = `Bearer ${newAccessToken}`;
        return api(originalRequest); 
      } catch (err) {
        store.dispatch(clearAccessToken());
        window.location.href = "/login";
      }
    }

    return Promise.reject(error);
  }
);

export default api;