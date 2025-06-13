import { createSlice, type PayloadAction } from "@reduxjs/toolkit";

export interface ScoreState {
  delegate_id: number;
  parameter_id: number;
  valueToDisplay: number;
}

const initialState = {
  scores: [] as ScoreState[],
}

const scoresSlice = createSlice({
  name: "scores",
  initialState,
  reducers: {
    setScores: (state, action: PayloadAction<ScoreState[]>) => {
      state.scores = action.payload;
    }
  }
})

export const { setScores } = scoresSlice.actions;

export default scoresSlice.reducer;