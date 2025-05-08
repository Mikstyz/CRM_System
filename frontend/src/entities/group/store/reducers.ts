import { GroupsState } from "@/entities/group/store/initialState.ts";

export const groupsReducer = {
  clearErrors: (state: GroupsState) => {
    state.error = undefined;
  },
};
