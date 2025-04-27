import {parseDateNameGroup} from "@/shared/lib/helpers/parseDateNameGroup.ts";
import {AppDispatch} from "@/app/store";
import {Id} from "@/shared/types";
import {updateGroupsThunks} from "@/entities/group/store/thunks.ts";
import {Group} from "@/entities/group/types";

interface Props {
  dispatch: AppDispatch;
  group: Group
  value: string;
}

export const handleTitleGroupSave = async ({dispatch, group, value}: Props) => {
  const resParseDateName = parseDateNameGroup(value);
  if (resParseDateName.status === "fulfilled") {
    dispatch(
      updateGroupsThunks({
        ...group,
        dateNameGroup: resParseDateName.dateNameGroup
      }),
    );
    console.log(`Group[${group.id}] → ${value}, `, { resParseDateName });
  } else {
    console.log("errors", resParseDateName.errors);
  }
}