import { configureStore } from '@reduxjs/toolkit';
// import counterReducer from '@/features/counter/counterSlice';   // add slices here

export const store = configureStore({
  reducer: {
    // counter: counterReducer,
  },
});

//— Type helpers ————————————————————————————————
export type RootState = ReturnType<typeof store.getState>;
export type AppDispatch = typeof store.dispatch;
