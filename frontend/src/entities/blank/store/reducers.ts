import { PayloadAction } from "@reduxjs/toolkit";
import { Id } from "@/shared/types";
import {
  blankInitialState,
  BlankState,
} from "@/entities/blank/store/initialState.ts";

export const blankReducers = {
  openBlank(state: BlankState, { payload }: PayloadAction<Id>) {
    state.isOpen = true;
    state.groupId = payload;
  },
  closeBlank(state: BlankState) {
    Object.assign(state, blankInitialState);
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
};
