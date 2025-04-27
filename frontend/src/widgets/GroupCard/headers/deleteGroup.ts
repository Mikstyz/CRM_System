import { AppDispatch } from "@/app/store";
import { Id } from "@/shared/types";
import { deleteGroupsThunks } from "@/entities/group/store/thunks.ts";

interface Props {
  dispatch: AppDispatch;
  groupId: Id;
}

export const headerDeleteGroup = async ({ dispatch, groupId }: Props) => {
  const isConfirmed = confirm("Вы уверены, что хотите удалить группу?");
  if (isConfirmed) {
    dispatch(deleteGroupsThunks(groupId));
  }
};
