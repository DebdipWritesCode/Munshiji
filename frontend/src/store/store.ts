import { configureStore } from '@reduxjs/toolkit';
import authSlice from '../slices/authSlice';
import allSheetsSlice from '../slices/allSheetsSlice';
import sheetDetailsSlice from '../slices/sheetDetailsSlice';
import scoresSlice from '../slices/scoresSlice';

export const store = configureStore({
  reducer: {
    auth: authSlice,
    allSheets: allSheetsSlice,
    sheetDetails: sheetDetailsSlice,
    scores: scoresSlice,
  },
});

//— Type helpers ————————————————————————————————
export type RootState = ReturnType<typeof store.getState>;
export type AppDispatch = typeof store.dispatch;
