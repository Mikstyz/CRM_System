import { groupReducer, sliceActions } from "@/entities/group/store/slice.ts";

export const { clearErrors, clearMessage } = sliceActions;
export const groupsActions = sliceActions;
export { groupReducer };
export type { GroupsState } from "@/entities/group/store/initialState";
