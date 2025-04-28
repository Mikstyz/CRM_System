import { blankReducer, sliceActions } from "./slice.ts";

export const {
  clearErrors: blankClearErrors,
  openBlank,
  closeBlank,
  setStudent,
} = sliceActions;

export { blankReducer };
export type { BlankState } from "./initialState";
