import { Id } from "@/shared/types";

export interface Student {
  id: Id;
  fullName: string;
  company?: string;
  startDateWork?: string;
  position?: string;
}
