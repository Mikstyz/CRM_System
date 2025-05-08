import { ActionReducerMapBuilder } from "@reduxjs/toolkit";
import { BlankState } from "@/entities/blank/store/initialState.ts";
import {
  handlePending,
  handleRejected,
} from "@/shared/lib/helpers/StoreHandlers.ts";
import {
  createStudentThunks,
  deleteStudentThunks,
  getAllStudentGroupThunks,
  updateStudentThunks,
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

  // CreateStudent
  builder
    .addCase(createStudentThunks.pending, handlePending<BlankState>)
    .addCase(createStudentThunks.fulfilled, (state, action) => {
      state.loading = false;
      state.studentsData.push(action.payload);
    })
    .addCase(createStudentThunks.rejected, handleRejected<BlankState>);

  // updateStudent
  builder
    .addCase(updateStudentThunks.pending, handlePending<BlankState>)
    .addCase(updateStudentThunks.fulfilled, (state, action) => {
      state.loading = false;
      state.selectStudent = action.payload;
      state.studentsData = state.studentsData.map((student) =>
        student.id === action.payload.id ? action.payload : student,
      );
    })
    .addCase(updateStudentThunks.rejected, handleRejected<BlankState>);

  // DeleteStudent
  builder
    .addCase(deleteStudentThunks.pending, handlePending<BlankState>)
    .addCase(deleteStudentThunks.fulfilled, (state, action) => {
      state.loading = false;
      state.studentsData = state.studentsData.filter(
        (student) => student.id !== action.payload,
      );
    })
    .addCase(deleteStudentThunks.rejected, handleRejected<BlankState>);
};
