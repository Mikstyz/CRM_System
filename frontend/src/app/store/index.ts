import { blankReducer } from "@/entities/blank/store/blankSlice";
import { groupFiltersReducer } from "@/features/FilterGroup/store/groupFiltersSlice.ts";
import { groupReducer } from "@/entities/group/store/slice.ts";
import { combineReducers, configureStore } from "@reduxjs/toolkit";

const rootReducer = combineReducers({
  groups: groupReducer,
  groupFilters: groupFiltersReducer,
  blank: blankReducer,
});

export const store = configureStore({
  reducer: rootReducer,
});

/* ─── typed hooks ───────────────────────────────── */
export type RootState = ReturnType<typeof store.getState>;
export type AppDispatch = typeof store.dispatch;
