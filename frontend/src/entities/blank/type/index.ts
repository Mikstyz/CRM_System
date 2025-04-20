import { Student } from "@/entities/student/types";
import { Id } from "@/shared/types";

export interface BlankState {
  isOpen: boolean;
  groupId: Id | null;
  studentId: Id | null;
  semester: number;
  company?: string | null;
  startDate: string; // YYYY‑MM‑DD
  position: string;
  studentsData: Student[];
}
