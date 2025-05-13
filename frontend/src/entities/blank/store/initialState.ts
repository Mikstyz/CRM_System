import { Id } from "@/shared/types";
import { Student } from "@/entities/student/types";
import { Loading } from "@/shared/types/store.type.ts";
import { Semester } from "@/entities/discipline/types";

export interface BlankState {
  isOpen: boolean;
  groupId: Id;
  semester: Semester;
  selectStudent?: Student;
  studentsData: Student[];
  loading: Loading;
  error?: string;
}

export const blankInitialState: BlankState = {
  isOpen: false,
  groupId: undefined,
  semester: "1",
  selectStudent: undefined,
  studentsData: [],
  loading: {
    status: "idle",
    message: undefined,
  },
  error: undefined,
};
