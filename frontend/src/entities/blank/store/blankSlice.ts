import { createSlice, PayloadAction } from "@reduxjs/toolkit";
import { RootState } from "@/app/store";
import { BlankState } from "../type";

const initialState: BlankState = {
  isOpen: false,
  groupId: null,
  studentId: null,
  semester: 1,
  company: "",
  startDate: "",
  position: "",
  studentsData: [],
};

const blankSlice = createSlice({
  name: "blank",
  initialState,
  reducers: {
    openBlank(state, { payload }: PayloadAction<string>) {
      state.isOpen = true;
      state.groupId = payload;
    },
    closeBlank(state) {
      Object.assign(state, initialState);
    },
  },
});

export const { openBlank, closeBlank } = blankSlice.actions;

export const blankReducer = blankSlice.reducer;

/* ---------- selectors ---------- */
export const selectBlank = (s: RootState) => s.blank;
export const selectGroupId = (s: RootState) => s.blank?.groupId;
