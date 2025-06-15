import { generatePdfThunks } from "@/entities/blank/store";
import { AppDispatch } from "@/app/store";
import { Group } from "@/entities/group/types";
import { Student } from "@/entities/student/types";
import { Semester } from "@/entities/discipline/types";

interface handleGetStudentProps {
  dispatch: AppDispatch;
  group: Group;
  student: Student;
  semester: Semester;
}
export const handleGetStudent = ({
  dispatch,
  group,
  student,
  semester,
}: handleGetStudentProps) => {
  dispatch(
    generatePdfThunks({
      group,
      student,
      semester,
    }),
  );
};
