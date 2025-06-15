import { useEffect } from "react";
import { ModalWrapper } from "@/shared/ui/ModalWrapper";
import { BlankPage } from "@/widgets/BlankPage";
import { FilterGroup } from "@/features/FilterGroup";
import { useAppDispatch, useAppSelector } from "@/shared/lib/hooks/redux";
import { selectFilteredGroups } from "@/entities/group/selectors";
import { getGroupsThunks } from "@/entities/group/store/thunks.ts";
import { clearErrors } from "@/entities/group/store";
import { closeBlank } from "@/entities/blank/store";
import { ModalErrorBoundary } from "@/widgets/ModalErrorBoundary";
import { ListGroup } from "@/features/ListGroup";
import { selectBlank } from "@/entities/blank/store/selectors.ts";

export function PagesListGroup() {
  const dispatch = useAppDispatch();
  const { isOpen } = useAppSelector(selectBlank);
  const groups = useAppSelector(selectFilteredGroups);

  useEffect(() => {
    dispatch(clearErrors());
    dispatch(getGroupsThunks());
  }, [dispatch]);

  return (
    <>
      <header className="p-4 flex gap-2 items-center">
        <h1 className="text-xl font-bold">Управление группами</h1>
      </header>

      <main className="p-4 flex flex-col gap-4">
        <section className="flex flex-row gap-6">
          <FilterGroup groupsLength={groups.length} />
          {/*<UsageInformation />*/}
        </section>

        <ListGroup groups={groups} />
      </main>

      <ModalErrorBoundary>
        <ModalWrapper isOpen={isOpen} onClose={() => dispatch(closeBlank())}>
          <BlankPage />
        </ModalWrapper>
      </ModalErrorBoundary>
    </>
  );
}
