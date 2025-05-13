import { PanelGroupCard } from "@/features/PanelGroupCard";
import { EditableTitle } from "@/shared/ui/EditableTitle";
import { useState } from "react";
import { SemesterDisciplines } from "@/features/SemesterDisciplines";
import { useAppDispatch } from "@/shared/lib/hooks/redux";
import { Group } from "@/entities/group/types";
import { headerDeleteGroup } from "@/widgets/GroupCard/headers/deleteGroup.ts";
import { handleTitleGroupSave } from "@/shared/lib/headers/titleGroupSave.ts";
import { headerDuplicateGroup } from "@/widgets/GroupCard/headers/duplicateGroup.ts";
import { openBlank } from "@/entities/blank/store";
import { useConfirm } from "@/shared/ui/ConfirmDialog";

interface Props {
  group: Group;
}

export function GroupCard({ group }: Props) {
  const [isExpanded, setIsExpanded] = useState(false);
  const dispatch = useAppDispatch();
  const confirm = useConfirm();
  const [err, setErr] = useState<string | null>(null);

  return (
    <section className="border-2 rounded-xl p-4">
      <header className="flex items-center justify-between">
        <h2 className="text-lg font-semibold">
          Группа:
          <EditableTitle
            key={group.name}
            initialValue={group.name}
            onSave={async (value) => {
              const ok = await handleTitleGroupSave({ dispatch, group, value });
              setErr(ok ? null : "Неверный формат имени");
            }}
            className="ml-1 w-min"
            error={err || undefined}
          />
        </h2>

        <PanelGroupCard
          onToggleExpand={() => setIsExpanded((prev) => !prev)}
          onDeleteGroup={() =>
            headerDeleteGroup({ dispatch, confirm, groupId: group.id })
          }
          onDuplicateGroup={() =>
            headerDuplicateGroup({ dispatch, groupId: group.id })
          }
          onOpenBlank={() => {
            dispatch(openBlank(group.id));
          }}
          isExpanded={isExpanded}
        />
      </header>

      {isExpanded && (
        <div className="border-t pt-4 mt-2 space-y-4">
          {/*
          // TODO:FISH Сделать чтобы дисциплины открывались в модалтном окне, для удобства.
          <ModalErrorBoundary>
        <ModalWrapper isOpen={} onClose={() => {})}>
        </ModalWrapper>
      </ModalErrorBoundary>
          */}
          {/* 1‑й семестр */}
          <SemesterDisciplines
            semester={"1"}
            items={group.disciplines[1]}
            groupId={group.id}
          />

          {/* 2‑й семестр */}
          <SemesterDisciplines
            semester={"2"}
            items={group.disciplines[2]}
            groupId={group.id}
          />
        </div>
      )}
    </section>
  );
}
