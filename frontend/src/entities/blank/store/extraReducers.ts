import { ActionReducerMapBuilder } from "@reduxjs/toolkit";
import { BlankState } from "@/entities/blank/store/initialState.ts";
import {
  handlePending,
  handleRejected,
} from "@/shared/lib/helpers/StoreHandlers.ts";
import {
  getAllStudentGroupThunks,
  getInfStudentIDThunks,
} from "@/entities/blank/store/thunks.ts";

export const blankExtraReducers = (
  builder: ActionReducerMapBuilder<BlankState>,
) => {
  // ========= Student =========

  // GET AllStudent
  builder
    .addCase(getAllStudentGroupThunks.pending, handlePending<BlankState>)
    .addCase(getAllStudentGroupThunks.fulfilled, (state, action) => {
      state.loading = false;
      state.studentsData = action.payload;
    })
    .addCase(getAllStudentGroupThunks.rejected, handleRejected<BlankState>);

  // GET InfStudentID
  builder
    .addCase(getInfStudentIDThunks.pending, handlePending<BlankState>)
    .addCase(getInfStudentIDThunks.fulfilled, (state, action) => {
      state.loading = false;
      state.studentsData = action.payload;
    })
    .addCase(getInfStudentIDThunks.rejected, handleRejected<BlankState>);

  // CreateStudent
  // builder
  //   .addCase(createStudentThunks.pending, handlePending<BlankState>)
  //   .addCase(createStudentThunks.fulfilled, (state, action) => {
  //     state.loading = false;
  //     state.studentsData = action.payload
  //   })
  //   .addCase(createStudentThunks.rejected, handleRejected<BlankState>);

  // // GET AllStudent
  // builder
  //   .addCase(getAllStudentGroupThunks.pending, handlePending<BlankState>)
  //   .addCase(getAllStudentGroupThunks.fulfilled, (state, action) => {
  //     state.loading = false;
  //     state.studentsData = action.payload
  //   })
  //   .addCase(getAllStudentGroupThunks.rejected, handleRejected<BlankState>);
  //
  // // GET AllStudent
  // builder
  //   .addCase(getAllStudentGroupThunks.pending, handlePending<BlankState>)
  //   .addCase(getAllStudentGroupThunks.fulfilled, (state, action) => {
  //     state.loading = false;
  //     state.studentsData = action.payload
  //   })
  //   .addCase(getAllStudentGroupThunks.rejected, handleRejected<BlankState>);
  //
  // // GET AllStudent
  // builder
  //   .addCase(getAllStudentGroupThunks.pending, handlePending<BlankState>)
  //   .addCase(getAllStudentGroupThunks.fulfilled, (state, action) => {
  //     state.loading = false;
  //     state.studentsData = action.payload
  //   })
  //   .addCase(getAllStudentGroupThunks.rejected, handleRejected<BlankState>);
  //
  // // GET AllStudent
  // builder
  //   .addCase(getAllStudentGroupThunks.pending, handlePending<BlankState>)
  //   .addCase(getAllStudentGroupThunks.fulfilled, (state, action) => {
  //     state.loading = false;
  //     state.studentsData = action.payload
  //   })
  //   .addCase(getAllStudentGroupThunks.rejected, handleRejected<BlankState>);
  //
  // // GET AllStudent
  // builder
  //   .addCase(getAllStudentGroupThunks.pending, handlePending<BlankState>)
  //   .addCase(getAllStudentGroupThunks.fulfilled, (state, action) => {
  //     state.loading = false;
  //     state.studentsData = action.payload
  //   })
  //   .addCase(getAllStudentGroupThunks.rejected, handleRejected<BlankState>);
  //
  // // GET AllStudent
  // builder
  //   .addCase(getAllStudentGroupThunks.pending, handlePending<BlankState>)
  //   .addCase(getAllStudentGroupThunks.fulfilled, (state, action) => {
  //     state.loading = false;
  //     state.studentsData = action.payload
  //   })
  //   .addCase(getAllStudentGroupThunks.rejected, handleRejected<BlankState>);
};
