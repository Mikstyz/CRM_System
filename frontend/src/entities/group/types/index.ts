import { Discipline, Semester } from "@/entities/discipline/types";
import { Id } from "@/shared/types";

export interface Group {
  id: Id;
  name: string;
  dateNameGroup: DateNameGroup;
  disciplines: Record<Semester, Discipline[]>;
}

export interface DateNameGroup {
  course: string;
  specialty: string;
  graduates: "9" | "11";
  groupNumber: number;
}
