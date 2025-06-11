import { createSlice, type PayloadAction } from "@reduxjs/toolkit";

interface ScoreSheetState {
  id: number | null;
  name: string | null;
  committee_name: string | null;
  chair: string | null;
  vice_chair?: string | null;
  rapporteur?: string | null;
  created_by: number | null;
  created_at: string | null;
  updated_at: string | null;
}

const initialState = {
  sheets: [] as ScoreSheetState[],
  loading: false,
};

const allSheetsSlice = createSlice({
  name: "allSheets",
  initialState,
  reducers: {
    setAllSheets: (state, action: PayloadAction<ScoreSheetState[]>) => {
      state.sheets = action.payload;
      state.loading = false;
    },
    clearAllSheets: () => {
      return initialState;
    },
    setLoading: (state, action: PayloadAction<boolean>) => {
      state.loading = action.payload;
    },
  },
});

export const { setAllSheets, clearAllSheets, setLoading } = allSheetsSlice.actions;
export default allSheetsSlice.reducer;