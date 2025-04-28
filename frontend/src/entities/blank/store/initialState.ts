import { Id } from "@/shared/types";
import { Student } from "@/entities/student/types";

export interface BlankState {
  isOpen: boolean;
  groupId: Id | null;
  studentId: Id | null;
  studentName: string;
  semester: number;
  company: string | null;
  startDate: string;
  position: string;
  studentsData: Student[];
  error: string | null;
  loading: boolean;
}

export const blankInitialState: BlankState = {
  isOpen: false,
  groupId: null,
  studentId: null,
  studentName: "",
  semester: 1,
  company: null,
  startDate: new Date().toISOString().slice(0, 10),
  position: "",
  studentsData: [],
  error: null,
  loading: false,
};
