import { Button } from "./components/ui/button"
import { useEffect } from "react";
import { useDispatch } from "react-redux";
import axios from "axios";
import { setAccessToken, clearAccessToken } from "./slices/authSlice";

const App = () => {
  const dispatch = useDispatch();

  useEffect(() => {
    const refreshToken = async () => {
      try {
        const res = await axios.post("/refresh_access_token", {});
        dispatch(setAccessToken(res.data.access_token));
      } catch {
        dispatch(clearAccessToken());
      }
    };
    refreshToken();
  }, [dispatch]);

  return (
    <div className=" ">
      <Button variant="default">Click Me</Button>
    </div>
  )
}

export default App