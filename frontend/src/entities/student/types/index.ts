export interface Student {
  id: string;
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
