import { z } from "zod";

export const blankSchema = z.object({
  semester: z.enum(["1", "2"], {
    errorMap: () => ({ message: "Допустимо только '1' или '2'" }),
  }),
  studentName: z.string().min(1, "Обязательное поле"),
  company: z.string().optional(),
  startDate: z.string().optional(),
  position: z.string().min(2),
});

export type FormValuesBlank = z.infer<typeof blankSchema>;
export type BlankSchema = z.output<typeof blankSchema>;
