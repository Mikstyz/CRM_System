import { Group } from "@/entities/group/types";

export interface GroupsState {
  list: Group[];
  loading: boolean;
  error: string | null;
}

export const groupInitialState: GroupsState = {
  list: [],
  error: null,
  loading: false,
};
