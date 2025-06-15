import { PayloadAction } from "@reduxjs/toolkit";
import { Id } from "@/shared/types";
import { BlankState } from "@/entities/blank/store/initialState.ts";
import { Group } from "@/entities/group/types";

export const blankReducers = {
  openBlank(state: BlankState, { payload }: PayloadAction<Group>) {
    state.isOpen = true;
    state.group = payload;
  },
  closeBlank(state: BlankState) {
    state.isOpen = false;
    state.group = undefined;
    state.selectStudent = undefined;
    state.error = undefined;
    state.studentsData = [];
  },
  setStudent(
    state: BlankState,
    { payload }: PayloadAction<{ id: Id; fullName: string }>,
  ) {
    if (!state.selectStudent) return;
    state.selectStudent.id = payload.id;
    state.selectStudent.fullName = payload.fullName;
  },
  clearErrors(state: BlankState) {
    state.error = undefined;
  },
  clearMessage(state: BlankState) {
    state.loading.message = undefined;
  },
};
