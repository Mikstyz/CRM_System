import { createSlice, PayloadAction } from "@reduxjs/toolkit";
import { Course, Graduates } from "@/entities/group/types";

export interface GroupFilters {
  course?: Course | null;
  specialty?: string | null;
  graduates?: Graduates | null;
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
