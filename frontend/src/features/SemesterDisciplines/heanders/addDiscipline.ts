import { Semester } from "@/entities/discipline/types";
import { Id } from "@/shared/types";
import { AppDispatch } from "@/app/store";
import { addDisciplinesThunks } from "@/entities/group/store/thunks.ts";

interface Props {
  dispatch: AppDispatch;
  groupId: Id;
  semester: Semester;
  newTitle: string;
}

export const headerAddDiscipline = ({
  dispatch,
  groupId,
  semester,
  newTitle,
}: Props) => {
  dispatch(addDisciplinesThunks({ groupId, semester, newTitle }));
};
