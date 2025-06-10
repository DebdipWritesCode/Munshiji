import { useEffect } from "react";
import { useDispatch, useSelector } from "react-redux";
import axios from "./api/axios";
import { setAccessToken, clearAccessToken } from "./slices/authSlice";
import type { RootState } from "./store/store";
import Router from "./router/Router";

const App = () => {
  const dispatch = useDispatch();
  const loading = useSelector((state: RootState) => state.auth.loading);

  useEffect(() => {
    const refreshToken = async () => {
      try {
        const res = await axios.post("/refresh_access_token", {}, { withCredentials: true }); // ✅ Ensure this
        dispatch(setAccessToken(res.data.access_token));
      } catch {
        dispatch(clearAccessToken());
      }
    };
    refreshToken();
  }, [dispatch]);

  if (loading) return <div>Loading...</div>; // 👈 Don't render anything until we know

  return <Router />;
};

export default App;
