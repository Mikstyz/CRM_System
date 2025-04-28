import { createSlice } from "@reduxjs/toolkit";
import { blankInitialState } from "./initialState.ts";
import { blankExtraReducers } from "./extraReducers.ts";
import { blankReducers } from "./reducers";

const blankSlice = createSlice({
  name: "blank",
  initialState: blankInitialState,
  reducers: blankReducers,
  extraReducers: blankExtraReducers,
});

export const sliceActions = blankSlice.actions;

export const blankReducer = blankSlice.reducer;
