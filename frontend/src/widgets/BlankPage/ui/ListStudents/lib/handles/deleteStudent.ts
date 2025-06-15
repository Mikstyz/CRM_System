import { deleteStudentThunks } from "@/entities/blank/store";
import { Id } from "@/shared/types";
import { AppDispatch } from "@/app/store";

interface handleDeleteProps {
  id: Id;
  dispatch: AppDispatch;
}
export const handleDelete = ({ id, dispatch }: handleDeleteProps) => {
  dispatch(deleteStudentThunks(id));
};
