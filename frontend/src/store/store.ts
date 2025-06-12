import { configureStore } from '@reduxjs/toolkit';
import authSlice from '../slices/authSlice';
import allSheetsSlice from '../slices/allSheetsSlice';
import sheetDetailsSlice from '../slices/sheetDetailsSlice';

export const store = configureStore({
  reducer: {
    auth: authSlice,
    allSheets: allSheetsSlice,
    sheetDetails: sheetDetailsSlice,
  },
});

//— Type helpers ————————————————————————————————
export type RootState = ReturnType<typeof store.getState>;
export type AppDispatch = typeof store.dispatch;
