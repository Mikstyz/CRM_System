import { useEffect } from "react";
import { GroupCard } from "../../widgets/GroupCard";
import { LinkDocument } from "@/shared/ui/LinkDocument";
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
  toggleExpand,
  deleteGroup as delGroupAction,
  addDiscipline,
  deleteDiscipline,
} from "@/entities/group/store/groupSlice";
import { useAppDispatch, useAppSelector } from "@/shared/lib/hooks/redux";
import { initialGroups } from "@/app/test/mock.test";
import { Semester } from "@/entities/discipline/types";
import { Id } from "@/shared/types";
import { selectFilteredGroups } from "@/entities/group/selectors";
import { Group } from "@/entities/group/types";
import { closeBlank, selectBlank } from "@/entities/blank/store/blankSlice";

export function PagesListGroup() {
  const dispatch = useAppDispatch();
  const { isOpen, groupId } = useAppSelector(selectBlank);
  const groups = useAppSelector(selectFilteredGroups);

  const currentGroupName = groups.find((g) => g.id === groupId)?.name ?? "";

  useEffect(() => {
    dispatch(setGroups(initialGroups));
  }, [dispatch]);

  return (
    <>
      <header className="p-4 flex gap-2 items-center">
        <h1 className="text-xl font-bold">Управление группами</h1>
        <LinkDocument href="#">Открыть XLSX групп</LinkDocument>
      </header>

      <main className="p-4 space-y-4">
        <FilterGroup groupsLength={groups.length} />

        {groups.length > 0 && (
          groups.map((group: Group) => (
            <GroupCard
              key={group.id}
              group={group}
              onToggleExpand={() => dispatch(toggleExpand(group.id))}
              onAddDiscipline={(semester: Semester) =>
                dispatch(
                  addDiscipline({
                    groupId: group.id,
                    semester,
                    disc: {
                      id: crypto.randomUUID(),
                      title: "Новая дисциплина",
                    },
                  }),
                )
              }
              onDeleteDiscipline={(semester: Semester, discId: Id) =>
                dispatch(
                  deleteDiscipline({ groupId: group.id, semester, discId }),
                )
              }
              onDeleteGroup={() => dispatch(delGroupAction(group.id))}
            />
          ))
        )}
      </main>

      <ModalErrorBoundary>
        <ModalWrapper isOpen={isOpen} onClose={() => dispatch(closeBlank())}>
          <BlankPage groupId={groupId} groupName={currentGroupName} />
        </ModalWrapper>
      </ModalErrorBoundary>
    </>
  );
}
