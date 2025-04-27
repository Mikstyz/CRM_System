import { AppDispatch } from "@/app/store";
import { Id } from "@/shared/types";
import { Semester } from "@/entities/discipline/types";
import { deleteDisciplinesThunks } from "@/entities/group/store/thunks.ts";

interface Props {
  dispatch: AppDispatch;
  groupId: Id;
  semester: Semester;
  discId: Id;
}
export const headerDeleteDiscipline = async ({
  dispatch,
  groupId,
  semester,
  discId,
}: Props) => {
  dispatch(deleteDisciplinesThunks({ groupId, semester, discId }));
};
