import { EditableTitle } from "src/shared/ui/EditableTitle";
import { Discipline } from "@/entities/discipline/types";

interface DisciplineItemProps {
  discipline: Discipline;
  onDelete: () => void;
  handleTitleSave: (newValue: string) => void;
}

export function DisciplineItem({
  discipline,
  onDelete,
  handleTitleSave,
}: DisciplineItemProps) {
  return (
    <div className="border rounded-lg p-2 flex justify-between items-center">
      <EditableTitle value={discipline.title} onSave={handleTitleSave} />
      <div>{}</div>
      <button
        onClick={onDelete}
        className="bg-red-400 hover:bg-red-500 text-white font-bold py-1 px-2 rounded"
        title="Удалить дисциплину"
      >
        ✕
      </button>
    </div>
  );
}
