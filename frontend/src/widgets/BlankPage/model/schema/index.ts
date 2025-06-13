import { Semester } from "@/entities/discipline/types";
import { z } from "zod";

// Updated preprocessing function with proper typing
const semesterEnum = ["1", "2"] as const;
const toSemesterOrUndef = (v: unknown): Semester | undefined => {
  if (typeof v !== "string") return undefined;
  return semesterEnum.includes(v as Semester) ? (v as Semester) : undefined;
};

export const blankSchema = z.object({
  semester: z
    .custom<Semester | undefined>(toSemesterOrUndef, {
      message: "Invalid semester",
    })
    .optional(),
  studentName: z.string().min(1, "Обязательное поле"),
  company: z.string().optional(),
  startDate: z.string().optional(),
  position: z.string().min(2),
});

export type FormValuesBlank = z.infer<typeof blankSchema>;
export type BlankSchema = z.output<typeof blankSchema>;
