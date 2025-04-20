import { Student } from "@/entities/student/types";
import { Id } from "@/shared/types";

export interface BlankState {
  isOpen: boolean;
  groupId: Id | null;
  studentId: Id | null; // 🆕
  studentName: string; // удобно для префилла
  semester: number;
  company: string | null;
  startDate: string;
  position: string;
  studentsData: Student[]; // как было
}
