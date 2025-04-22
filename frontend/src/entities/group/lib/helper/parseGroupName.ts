export interface ParsedName {
  course: number;
  specialty: string;
  graduates: "9" | "11";
  groupNumber: number;
}

export const parseGroupName = (name: string): ParsedName | null => {
  // 1ИСП11‑45  →  [1] [ИСП] [11] [45]
  const match = name.match(/^(\d)([А-ЯA-Z]+)(\d{1,2})[‑-](\d{1,})$/i);
  if (!match) return null;
  return {
    course: Number(match[1]),
    specialty: match[2].toUpperCase(),
    graduates: match[3] as "9" | "11",
    groupNumber: Number(match[4]),
  };
};
