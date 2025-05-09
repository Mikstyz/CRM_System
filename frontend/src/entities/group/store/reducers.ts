import { GroupsState } from "@/entities/group/store/initialState.ts";
import { PayloadAction } from "@reduxjs/toolkit";

export const groupsReducer = {
  clearErrors: (state: GroupsState) => {
    state.error = undefined;
  },
  setError: (state: GroupsState, { payload }: PayloadAction<string>) => {
    state.error = payload;
  },
};
