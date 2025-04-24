import { Id } from "@/shared/types";

export interface Student {
  id: Id;
  fullName: string;
  groups: string[];
  semester: number;
  company?: string | null;
  startDateWork: Date;
  position: string;
}

export interface UserState {
  listStudents: Student[];
}
