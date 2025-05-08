import { Group } from "@/entities/group/types";

export interface GroupsState {
  list: Group[];
  loading: boolean;
  error: string | undefined;
}

export const groupInitialState: GroupsState = {
  list: [],
  error: undefined,
  loading: false,
};
