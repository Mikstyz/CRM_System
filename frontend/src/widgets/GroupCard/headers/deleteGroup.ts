import { AppDispatch } from "@/app/store";
import { Id } from "@/shared/types";
import { deleteGroupsThunks } from "@/entities/group/store/thunks.ts";
import { ConfirmOptions } from "@/shared/ui/ConfirmDialog";

interface Props {
  dispatch: AppDispatch;
  groupId: Id;
  confirm: (d: ConfirmOptions) => unknown;
}

export const headerDeleteGroup = async ({
  dispatch,
  groupId,
  confirm,
}: Props) => {
  // const isConfirmed = confirm("Вы уверены, что хотите удалить группу?");
  const ok = await confirm({
    title: "Удалить группу?",
    description: "Это действие нельзя отменить.",
    confirmText: "Да, удалить",
    cancelText: "Отмена",
  });
  if (ok) dispatch(deleteGroupsThunks(groupId));
};
