import { RootState } from "@/app/store";

export const selectBlank = (s: RootState) => s.blank;
export const selectGroupId = (s: RootState) => s.blank?.groupId;
