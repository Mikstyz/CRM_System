import { createSelector } from "@reduxjs/toolkit";
import { RootState } from "@/app/store";
import { parseGroupName } from "./lib/helper/parseGroupName";
import { Group } from "./types";

export const selectFilteredGroups = createSelector(
  [(s: RootState) => s.groups.list, (s: RootState) => s.groupFilters],
  (groups, f) =>
    groups.filter((g: Group) => {
      const meta = parseGroupName(g.name);
      if (!meta) return false;

      return (
        (f.course ? String(meta.course) === f.course : true) &&
        (f.specialty
          ? meta.specialty.startsWith(f.specialty.toUpperCase())
          : true) &&
        (f.graduates ? meta.graduates === f.graduates : true) &&
        (f.groupNumber ? meta.groupNumber === f.groupNumber : true)
      );
    }),
);
