import {AppDispatch} from "@/app/store";
import {Id} from "@/shared/types";
import {Semester} from "@/entities/discipline/types";
import {updateDisciplinesThunks} from "@/entities/group/store/thunks.ts";

interface Props {
  dispatch: AppDispatch;
  groupId: Id;
  semester: Semester;
  newTitle: string;
  discId: Id
}

export const headerUpdateDiscipline = ({
                                    dispatch,
                                    groupId,
                                    semester,
                                    newTitle,
                                    discId
                                  }: Props) => {
  dispatch(updateDisciplinesThunks({groupId, semester, newTitle, discId}));
};