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
    addNewScore: (state, action) => {
      state.scores = [...state.scores, action.payload];
    },
    updateScore: (state, action) => {
      const { id, value } = action.payload;
      const score = state.scores.find((score) => score.id === id);
      if (score) {
        score.value = value;
      }
    },
    updateScoreNote: (state, action) => {
      const { id, note } = action.payload;
      const score = state.scores.find((score) => score.id === id);
      if (score) {
        score.note = note;
      }
    },
    deleteScore: (state, action) => {
      const id = action.payload;
      state.scores = state.scores.filter((score) => score.id !== id);
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
  addNewScore,
  setScores,
  updateScore,
  deleteScore,
  updateScoreNote,
  setLoading,
  setError,
} = sheetDetails.actions;

export default sheetDetails.reducer;
