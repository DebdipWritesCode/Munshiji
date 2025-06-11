import { configureStore } from '@reduxjs/toolkit';
import authSlice from '../slices/authSlice';
import allSheetsSlice from '../slices/allSheetsSlice';

export const store = configureStore({
  reducer: {
    auth: authSlice,
    allSheets: allSheetsSlice,
  },
});

//— Type helpers ————————————————————————————————
export type RootState = ReturnType<typeof store.getState>;
export type AppDispatch = typeof store.dispatch;
