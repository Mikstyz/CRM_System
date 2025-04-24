import { useEffect, useState } from "react";
import { GroupCard } from "@/widgets/GroupCard";
import { ModalWrapper } from "@/shared/ui/ModalWrapper";
import { BlankPage } from "../BlankPage";
import { FilterGroup } from "@/features/FilterGroup";
import { Component, ReactNode } from "react";

class ModalErrorBoundary extends Component<
  { children: ReactNode },
  { hasError: boolean }
> {
  state = { hasError: false };

  static getDerivedStateFromError() {
    return { hasError: true };
  }

  componentDidCatch(err: unknown, info: unknown) {
    console.error("Modal crashed:", err, info);
  }

  render() {
    return this.state.hasError ? (
      <p className="p-4 text-red-600">Что‑то пошло не так…</p>
    ) : (
      this.props.children
    );
  }
}

import {
  setGroups,
  deleteGroup as delGroupAction,
  addDiscipline,
  deleteDiscipline,
} from "@/entities/group/store/groupSlice";
import { useAppDispatch, useAppSelector } from "@/shared/lib/hooks/redux";
import { Semester } from "@/entities/discipline/types";
import { Id } from "@/shared/types";
import { selectFilteredGroups } from "@/entities/group/selectors";
import { Group } from "@/entities/group/types";
import { closeBlank, selectBlank } from "@/entities/blank/store/blankSlice";
import {
  Add_SubjectByGroupId,
  Delete_GroupById,
  Delete_SubjectById,
  Inf_AllGroup,
} from "@wails/go/main/App";
import { ApiResponse, CreateApiResponse } from "@/shared/types/api.type.ts";

export function PagesListGroup() {
  const dispatch = useAppDispatch();
  const { isOpen, groupId } = useAppSelector(selectBlank);
  const groups = useAppSelector(selectFilteredGroups);

  const currentGroupName = groups.find((g) => g.id === groupId)?.name ?? "";

  const [initialGroupsAPI, setInitialGroupsAPI] = useState<Group[] | []>([]);

  async function getGroupsAPI() {
    const res = await Inf_AllGroup();
    const GroupsAll = res.Groups;
    if (GroupsAll !== null && GroupsAll !== undefined) {
      const Groups = GroupsAll.map((g) => {
        return {
          id: g.Id,
          name: `${g.Well}${g.Speciality}${g.GClass}-${g.Number}`,
          dateNameGroup: {
            course: g.Well.toString(),
            specialty: g.Speciality,
            graduates: g.GClass.toString(),
            groupNumber: g.Number,
          },
          disciplines: {
            "1": [],
            "2": [],
          },
        };
      });
      setInitialGroupsAPI(Groups as Group[]);
    }

    return res.Groups;
  }

  useEffect(() => {
    getGroupsAPI();
  }, []);

  useEffect(() => {
    console.log("getGroupsAPI", initialGroupsAPI);
    dispatch(setGroups(initialGroupsAPI));
  }, [dispatch, initialGroupsAPI]);

  return (
    <>
      <header className="p-4 flex gap-2 items-center">
        <h1 className="text-xl font-bold">Управление группами</h1>
      </header>

      <main className="p-4 space-y-4">
        <FilterGroup groupsLength={groups.length} />

        {groups.length > 0 &&
          groups.map((group: Group) => (
            <GroupCard
              key={group.id}
              group={group}
              onAddDiscipline={async (semester: Semester) => {
                const res = (await Add_SubjectByGroupId(
                  group.id,
                  "Новая дисциплина",
                )) as CreateApiResponse;
                if (res.Id) {
                  dispatch(
                    addDiscipline({
                      groupId: group.id,
                      semester,
                      disc: {
                        id: res.Id,
                        title: "Новая дисциплина",
                      },
                    }),
                  );
                }
              }}
              onDeleteDiscipline={async (semester: Semester, discId: Id) => {
                const res = (await Delete_SubjectById(discId)) as ApiResponse;
                if (res.Code === 200) {
                  dispatch(
                    deleteDiscipline({
                      groupId: group.id,
                      semester,
                      discId,
                    }),
                  );
                }
              }}
              onDeleteGroup={async () => {
                const res = (await Delete_GroupById(group.id)) as ApiResponse;
                if (res.Code === 200) {
                  dispatch(delGroupAction(group.id));
                }
              }}
            />
          ))}
      </main>

      <ModalErrorBoundary>
        <ModalWrapper isOpen={isOpen} onClose={() => dispatch(closeBlank())}>
          <BlankPage groupId={groupId} groupName={currentGroupName} />
        </ModalWrapper>
      </ModalErrorBoundary>
    </>
  );
}
