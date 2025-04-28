import { ActionReducerMapBuilder } from "@reduxjs/toolkit";
import { BlankState } from "@/entities/blank/store/initialState.ts";

export const blankExtraReducers = (
  builder: ActionReducerMapBuilder<BlankState>,
) => {
  // ========= GROUPS =========
  // GET GROUPS
  // builder
  //   .addCase(getGroupsThunks.pending, handlePending<BlankState>)
  //   .addCase(getGroupsThunks.fulfilled, (state, action) => {
  //     state.loading = false;
  //   })
  //   .addCase(getGroupsThunks.rejected, handleRejected<BlankState>);
};
