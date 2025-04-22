import { createSlice, PayloadAction } from "@reduxjs/toolkit";

export interface GroupFilters {
  course?: number | null;
  specialty?: string | null;
  graduates?: "9" | "11" | null;
  groupNumber?: number | null;
}

const initialState: GroupFilters = {};

const groupFiltersSlice = createSlice({
  name: "groupFilters",
  initialState,
  reducers: {
    /** full replace (used on every keystroke) */
    setFilters: (_state, { payload }: PayloadAction<GroupFilters>) => payload,
    /** reset – e.g. on form reset */
    clearFilters: () => initialState,
  },
});

export const { setFilters, clearFilters } = groupFiltersSlice.actions;
export const groupFiltersReducer = groupFiltersSlice.reducer;
