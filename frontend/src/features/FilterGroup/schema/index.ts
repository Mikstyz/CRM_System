import { z } from "zod";

/* helper unchanged */
const toNumOrUndef = (v: unknown) => {
  const n = Number(v);
  return v === "" || v === null || Number.isNaN(n) ? undefined : n;
};

export const filterSchema = z.object({
  course: z.preprocess(toNumOrUndef, z.enum(["1", "2", "3", "4"]).optional()),
  specialty: z
    .string()
    .trim()
    .transform((s) => (s === "" ? undefined : s.toUpperCase()))
    .refine((v) => v === undefined || v.length >= 2, {
      message: "Минимум 2 символа",
    })
    .optional(),
  graduates: z.preprocess(
    (v) => (v === "" ? undefined : v),
    z.enum(["9", "11"]).optional(),
  ),
  groupNumber: z.preprocess(toNumOrUndef, z.number().int().optional()),
});

export type FiltersRaw = z.input<typeof filterSchema>; // unknowns allowed
export type Filters = z.output<typeof filterSchema>; // numbers | undefined
