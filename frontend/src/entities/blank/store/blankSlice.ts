import { createSlice, PayloadAction } from "@reduxjs/toolkit";
import { RootState } from "@/app/store";
import { BlankState } from "../type";
import { Id } from "@/shared/types";

const initialState: BlankState = {
  isOpen: false,
  groupId: null,
  studentId: null,
  studentName: "",
  semester: 1,
  company: null,
  startDate: new Date().toISOString().slice(0, 10),
  position: "",
  studentsData: [],
};

const blankSlice = createSlice({
  name: "blank",
  initialState,
  reducers: {
    openBlank(state, { payload }: PayloadAction<Id>) {
      state.isOpen = true;
      state.groupId = payload;
    },
    closeBlank(state) {
      Object.assign(state, initialState);
    },
    setStudent(
      state,
      { payload }: PayloadAction<{ id: Id; fullName: string }>,
    ) {
      state.studentId = payload.id;
      state.studentName = payload.fullName;
    },
  },
});

export const { setStudent, openBlank, closeBlank } = blankSlice.actions;

export const blankReducer = blankSlice.reducer;

/* ---------- selectors ---------- */
export const selectBlank = (s: RootState) => s.blank;
export const selectGroupId = (s: RootState) => s.blank?.groupId;
