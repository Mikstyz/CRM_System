import { DisciplineItem } from "../DisciplineItem";
import { PanelGroupCard } from "@/features/PanelGroupCard";
import { ListChildComponentProps } from "react-window";
import { EditableTitle } from "@/shared/ui/EditableTitle";
import { memo, useState } from "react";
import { SemesterDisciplines } from "@/features/SemesterDisciplines";
import { openBlank } from "@/entities/blank/store/blankSlice";
import { useAppDispatch } from "@/shared/lib/hooks/redux";
import { Group } from "@/entities/group/types";
import { headerDeleteGroup } from "@/widgets/GroupCard/headers/deleteGroup.ts";
import {handleTitleGroupSave} from "@/widgets/GroupCard/headers/titleGroupSave.ts";
import {headerDuplicateGroup} from "@/widgets/GroupCard/headers/duplicateGroup.ts";
import {Discipline} from "@/entities/discipline/types";

interface DisciplineItemProps {
  discipline: Discipline;
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
}

export function GroupCard({
  group,
}: Props) {
  const [isExpanded, setIsExpanded] = useState(false);
  const dispatch = useAppDispatch();

  return (
    <section className="border-2 rounded-xl p-4">
      <header className="flex items-center justify-between">
        <h2 className="text-lg font-semibold">
          Группа:
          <EditableTitle
            initialValue={group.name}
            onSave={(value) => handleTitleGroupSave({dispatch, group, value})}
            className="ml-1 w-min"
          />
        </h2>

        <PanelGroupCard
          onToggleExpand={() => setIsExpanded((prev) => !prev)}
          onDeleteGroup={() => headerDeleteGroup({ dispatch, groupId: group.id })}
          onDuplicateGroup={() => headerDuplicateGroup({dispatch, groupId: group.id})}
          onOpenBlank={() => {
            dispatch(openBlank(group.id));
          }}
          isExpanded={isExpanded}
        />
      </header>

      {isExpanded && (
        <div className="border-t pt-4 mt-2 space-y-4">
          {/* 1‑й семестр */}
          <SemesterDisciplines
            semester={1}
            items={group.disciplines[1]}
            groupId={group.id}
          />

          {/* 2‑й семестр */}
          <SemesterDisciplines
            semester={2}
            items={group.disciplines[2]}
            groupId={group.id}
          />
        </div>
      )}
    </section>
  );
}
