import { ActionReducerMapBuilder } from "@reduxjs/toolkit";
import { BlankState } from "@/entities/blank/store/initialState.ts";
import {
  handlePending,
  handleRejected,
} from "@/shared/lib/helpers/StoreHandlers.ts";
import {
  createStudentThunks,
  deleteStudentThunks,
  generatePdfThunks,
  getAllStudentGroupThunks,
  saveOrUpdateStudentThunks,
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
      state.loading = {
        status: "succeeded",
      };
      state.studentsData = action.payload;
    })
    .addCase(getAllStudentGroupThunks.rejected, handleRejected<BlankState>);

  // CreateStudent
  builder
    .addCase(createStudentThunks.pending, handlePending<BlankState>)
    .addCase(createStudentThunks.fulfilled, (state, action) => {
      state.loading = {
        status: "succeeded",
        message: "Успешно добавлен",
      };
      state.studentsData.push(action.payload);
    })
    .addCase(createStudentThunks.rejected, handleRejected<BlankState>);

  // updateStudent
  builder
    .addCase(updateStudentThunks.pending, handlePending<BlankState>)
    .addCase(updateStudentThunks.fulfilled, (state, action) => {
      state.loading = {
        status: "succeeded",
        message: "Успешно обновлен",
      };
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
      state.loading = {
        status: "succeeded",
        message: "Успешно удален",
      };
      state.studentsData = state.studentsData.filter(
        (student) => student.id !== action.payload,
      );
    })
    .addCase(deleteStudentThunks.rejected, handleRejected<BlankState>);

  builder
    .addCase(saveOrUpdateStudentThunks.pending, handlePending<BlankState>)
    .addCase(saveOrUpdateStudentThunks.fulfilled, (state, { payload }) => {
      state.loading = {
        status: "succeeded",
        message: "Успешно сохранино",
      };
      /* заменяем или добавляем */
      const idx = state.studentsData.findIndex((s) => s.id === payload.id);
      if (idx === -1) state.studentsData.push(payload);
      else state.studentsData[idx] = payload;
      state.selectStudent = payload;
    })
    .addCase(saveOrUpdateStudentThunks.rejected, handleRejected<BlankState>);

  /* PDF */
  builder
    .addCase(generatePdfThunks.pending, handlePending<BlankState>)
    .addCase(generatePdfThunks.fulfilled, (state) => {
      state.loading = {
        status: "succeeded",
        message: "Успешно сгенерирован",
      };
    })
    .addCase(generatePdfThunks.rejected, handleRejected<BlankState>);
};
