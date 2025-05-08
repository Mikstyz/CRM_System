import { Id } from "@/shared/types";

export interface Student {
  id: Id;
  fullName: string;
  company: string | undefined;
  startDateWork: Date | undefined;
  position: string | undefined;
}
