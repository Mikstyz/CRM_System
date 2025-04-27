import { AppDispatch } from "@/app/store";
import { Id } from "@/shared/types";
import { duplicateGroupThunks } from "@/entities/group/store/thunks.ts";

interface Props {
  dispatch: AppDispatch;
  groupId: Id;
}

export const headerDuplicateGroup = ({ dispatch, groupId }: Props) => {
  const isConfirmed = confirm("Вы уверены, что хотите дублировать группу?");
  if (isConfirmed) {
    dispatch(duplicateGroupThunks(groupId));
  }
};
