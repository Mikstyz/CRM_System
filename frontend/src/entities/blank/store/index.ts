import { blankReducer, sliceActions } from "./slice.ts";

export const {
  clearErrors: blankClearErrors,
  clearMessage: blankClearMessage,
  openBlank,
  closeBlank,
  setStudent,
} = sliceActions;

export { blankReducer };
export type { BlankState } from "./initialState";
