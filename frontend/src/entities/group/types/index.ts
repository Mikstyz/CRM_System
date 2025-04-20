import { Discipline, Semester } from "@/entities/discipline/types";

export interface Group {
  id: string;
  name: string;
  isExpanded: boolean;
  disciplines: Record<Semester, Discipline[]>;
}
