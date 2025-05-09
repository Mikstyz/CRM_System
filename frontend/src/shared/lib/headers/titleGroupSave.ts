import { parseDateNameGroup } from "@/shared/lib/helpers/parseDateNameGroup.ts";
import { AppDispatch } from "@/app/store";
import { updateGroupsThunks } from "@/entities/group/store/thunks.ts";
import { Group } from "@/entities/group/types";
import { clearFilters } from "@/features/FilterGroup/store/groupFiltersSlice.ts";
import { groupsActions } from "@/entities/group/store";

interface Props {
  dispatch: AppDispatch;
  group: Group;
  value: string;
}

export const handleTitleGroupSave = async ({
  dispatch,
  group,
  value,
}: Props) => {
  const resParseDateName = parseDateNameGroup(value);
  if (resParseDateName.status === "fulfilled") {
    dispatch(
      updateGroupsThunks({
        ...group,
        dateNameGroup: resParseDateName.dateNameGroup,
      }),
    );
    dispatch(clearFilters());
    return true;
  } else {
    dispatch(groupsActions.setError(resParseDateName.errors.join("; ")));
    return false;
  }
};
