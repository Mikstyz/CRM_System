import { createSlice } from "@reduxjs/toolkit";
import { groupsReducer } from "@/entities/group/store/reducers.ts";
import { groupInitialState } from "@/entities/group/store/initialState.ts";
import { groupsExtraReducers } from "@/entities/group/store/extraReducers.ts";

export const slice = createSlice({
  name: "groups",
  initialState: groupInitialState,
  reducers: groupsReducer,
  extraReducers: groupsExtraReducers,
});

export const sliceActions = slice.actions;

export const groupReducer = slice.reducer;
