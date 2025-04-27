import { Discipline, Semester } from "@/entities/discipline/types";
import { Id } from "@/shared/types";

export type Disciplines = Record<Semester, Discipline[]>;

export interface Group {
  id: Id;
  name: string;
  dateNameGroup: DateNameGroup;
  disciplines: Disciplines;
}
export type Course = "1" | "2" | "3" | "4";
export type Graduates = "9" | "11";

export interface DateNameGroup {
  course: Course;
  specialty: string; // минимум 2 буквы
  graduates: Graduates;
  groupNumber: number;
}
