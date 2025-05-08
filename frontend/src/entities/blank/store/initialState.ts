import { Id } from "@/shared/types";
import { Student } from "@/entities/student/types";

export interface BlankState {
  isOpen: boolean;
  groupId: Id | undefined;
  semester: "1" | "2" | undefined;
  selectStudent: Student | undefined;
  studentsData: Student[];
  error: string | undefined;
  loading: boolean;
}

export const blankInitialState: BlankState = {
  isOpen: false,
  groupId: undefined,
  semester: undefined,
  selectStudent: undefined,
  studentsData: [],
  error: undefined,
  loading: false,
};
