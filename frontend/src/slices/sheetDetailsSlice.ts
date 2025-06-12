import { createSlice } from "@reduxjs/toolkit";
import type {
  Score,
  ScoreSheet,
  Delegate,
  Parameter,
} from "../utils/getDataInterfaces";

interface SheetDetailsState {
  score_sheet: ScoreSheet | null;
  delegates: Delegate[];
  parameters: Parameter[];
  scores: Score[];
  loading: boolean;
  error: string | null;
}

const initialState: SheetDetailsState = {
  score_sheet: null,
  delegates: [],
  parameters: [],
  scores: [],
  loading: false,
  error: null,
};

const sheetDetails = createSlice({
  name: "sheetDetails",
  initialState,
  reducers: {
    setScoreSheet: (state, action) => {
      state.score_sheet = action.payload;
    },
    setDelegates: (state, action) => {
      state.delegates = action.payload;
    },
    setParameters: (state, action) => {
      state.parameters = action.payload;
    },
    setScores: (state, action) => {
      state.scores = action.payload;
    },
    setLoading: (state, action) => {
      state.loading = action.payload;
    },
    setError: (state, action) => {
      state.error = action.payload;
    },
  },
});

export const {
  setScoreSheet,
  setDelegates,
  setParameters,
  setScores,
  setLoading,
  setError,
} = sheetDetails.actions;

export default sheetDetails.reducer;
