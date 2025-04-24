import { DisciplineItem } from "../DisciplineItem";
import { PanelGroupCard } from "@/features/PanelGroupCard";
import { ListChildComponentProps } from "react-window";
import { EditableTitle } from "@/shared/ui/EditableTitle";
import { memo, useCallback, useState } from "react";
import { Semester } from "@/entities/discipline/types";
import { SemesterDisciplines } from "@/features/SemesterDisciplines";
import { Id } from "@/shared/types";
import { openBlank } from "@/entities/blank/store/blankSlice";
import { duplicateGroup } from "@/entities/group/store/groupSlice";
import { useAppDispatch } from "@/shared/lib/hooks/redux";
import { Group } from "@/entities/group/types";

interface DisciplineItemProps {
  discipline: any;
  onDelete: () => void;
}

const Row = memo(
  ({ data, index, style }: ListChildComponentProps<DisciplineItemProps[]>) => {
    const { discipline, onDelete } = data[index];
    return (
      <div style={style}>
        <DisciplineItem discipline={discipline} onDelete={onDelete} />
      </div>
    );
  },
);
Row.displayName = "Row";

interface Props {
  group: Group;
  onAddDiscipline: (semester: Semester) => void;
  onDeleteGroup: () => void;
  onDeleteDiscipline: (semester: Semester, id: Id) => void;
}

export function GroupCard({
  group: { id, name, disciplines },
  onAddDiscipline,
  onDeleteGroup,
  onDeleteDiscipline,
}: Props) {
  const [isExpanded, setIsExpanded] = useState(false);
  const dispatch = useAppDispatch();
  const handleTitleSave = useCallback(
    (value: string) => console.log(`Group[${id}] → ${value}`),
    [id],
  );

  return (
    <section className="border-2 rounded-xl p-4">
      <header className="flex items-center justify-between">
        <h2 className="text-lg font-semibold">
          Группа:
          <EditableTitle
            initialValue={name}
            onSave={handleTitleSave}
            className="ml-1 w-min"
          />
        </h2>

        <PanelGroupCard
          onToggleExpand={() => setIsExpanded((prev) => !prev)}
          onDeleteGroup={() => {
            const isConfirmed = confirm(
              "Вы уверены, что хотите удалить группу?",
            );
            if (isConfirmed) onDeleteGroup();
          }}
          onDuplicateGroup={() => dispatch(duplicateGroup(id))}
          onOpenBlank={() => dispatch(openBlank(id))}
          isExpanded={isExpanded}
        />
      </header>

      {isExpanded && (
        <div className="border-t pt-4 mt-2 space-y-4">
          {/* 1‑й семестр */}
          <SemesterDisciplines
            semester={1}
            items={disciplines[1]}
            onAdd={() => onAddDiscipline(1 as Semester)}
            onDelete={(id) => onDeleteDiscipline(1 as Semester, id)}
          />

          {/* 2‑й семестр */}
          <SemesterDisciplines
            semester={2}
            items={disciplines[2]}
            onAdd={() => onAddDiscipline(2 as Semester)}
            onDelete={(id) => onDeleteDiscipline(2 as Semester, id)}
          />
        </div>
      )}
    </section>
  );
}
