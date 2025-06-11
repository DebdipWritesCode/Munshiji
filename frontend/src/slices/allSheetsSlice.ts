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

const initialState: ScoreSheetState[] = [];

const allSheetsSlice = createSlice({
  name: "allSheets",
  initialState,
  reducers: {
    setAllSheets: (_, action: PayloadAction<ScoreSheetState[]>) => {
      return action.payload;
    },
    clearAllSheets: () => {
      return initialState;
    },
  },
});

export const { setAllSheets, clearAllSheets } = allSheetsSlice.actions;
export default allSheetsSlice.reducer;