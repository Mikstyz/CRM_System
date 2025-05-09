import { useEffect } from "react";
import { GroupCard } from "@/widgets/GroupCard";
import { ModalWrapper } from "@/shared/ui/ModalWrapper";
import { BlankPage } from "@/widgets/BlankPage";
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

import { useAppDispatch, useAppSelector } from "@/shared/lib/hooks/redux";
import { selectFilteredGroups } from "@/entities/group/selectors";
import { Group } from "@/entities/group/types";
import { getGroupsThunks } from "@/entities/group/store/thunks.ts";
import { clearErrors } from "@/entities/group/store";
import { selectBlank } from "@/entities/blank/store/selectors.ts";
import { closeBlank } from "@/entities/blank/store";
import { ErrorToast } from "@/shared/ui/ErrorToast";

export function PagesListGroup() {
  const dispatch = useAppDispatch();
  const { isOpen, groupId } = useAppSelector(selectBlank);
  const groups = useAppSelector(selectFilteredGroups);
  useEffect(() => {
    console.log("groups", groups);
  }, [groups]);
  const currentGroup = groups.find((g) => g.id === groupId);

  useEffect(() => {
    dispatch(clearErrors());
    dispatch(getGroupsThunks());
  }, [dispatch]);

  return (
    <>
      <header className="p-4 flex gap-2 items-center">
        <h1 className="text-xl font-bold">Управление группами</h1>
      </header>

      <main className="p-4 space-y-4">
        <FilterGroup groupsLength={groups.length} />

        {groups.length > 0 &&
          groups.map((group: Group) => (
            <GroupCard key={group.id} group={group} />
          ))}
      </main>

      <ModalErrorBoundary>
        <ModalWrapper isOpen={isOpen} onClose={() => dispatch(closeBlank())}>
          <BlankPage group={currentGroup} />
        </ModalWrapper>
      </ModalErrorBoundary>
      <ErrorToast />
    </>
  );
}
