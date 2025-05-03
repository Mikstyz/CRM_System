import { AppDispatch } from "@/app/store";
import { Id } from "@/shared/types";
import { duplicateGroupThunks } from "@/entities/group/store/thunks.ts";

interface Props {
  dispatch: AppDispatch;
  groupId: Id;
}

export const headerDuplicateGroup = ({ dispatch, groupId }: Props) => {
  dispatch(duplicateGroupThunks(groupId));
};
