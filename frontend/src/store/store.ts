import { configureStore } from '@reduxjs/toolkit';
import authSlice from '../slices/authSlice';

export const store = configureStore({
  reducer: {
    auth: authSlice,
  },
});

//— Type helpers ————————————————————————————————
export type RootState = ReturnType<typeof store.getState>;
export type AppDispatch = typeof store.dispatch;
