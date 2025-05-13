import { Group } from "@/entities/group/types";
import { Loading } from "@/shared/types/store.type.ts";

export interface GroupsState {
  list: Group[];
  loading: Loading;
  error?: string;
}

export const groupInitialState: GroupsState = {
  list: [],
  error: undefined,
  loading: {
    status: "idle",
    message: undefined,
  },
};
