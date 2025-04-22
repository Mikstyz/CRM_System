import { FixedSizeList } from "react-window";
import { Discipline, Semester } from "@/entities/discipline/types";
import { DisciplineItem } from "@/widgets/DisciplineItem";

interface Props {
  semester: Semester;
  items: Discipline[];
  onAdd(): void;
  onDelete(id: string): void;
}

export const SemesterDisciplines = ({
  semester,
  items,
  onAdd,
  onDelete,
}: Props) => {
  const itemData = items.map((d) => ({
    discipline: d,
    onDelete: () => onDelete(d.id),
  }));

  return (
    <details className="border rounded-xl">
      <summary className="cursor-pointer select-none py-1 px-3 bg-slate-100 rounded-xl">
        {semester}‑й семестр
      </summary>

      <div className="p-3 space-y-2">
        <button
          onClick={onAdd}
          className="bg-green-500 hover:bg-green-600 text-white font-bold py-1 px-3 rounded"
        >
          + предмет
        </button>

        {items.length === 0 ? (
          <p className="text-sm italic text-gray-500">Список пуст</p>
        ) : (
          <FixedSizeList
            height={Math.min(300, items.length * 55)}
            itemCount={items.length}
            itemSize={55}
            width="100%"
            itemData={itemData}
          >
            {({ style, index, data }) => {
              const { discipline, onDelete } = data[index];
              return (
                <div style={style}>
                  <DisciplineItem discipline={discipline} onDelete={onDelete} />
                </div>
              );
            }}
          </FixedSizeList>
        )}
      </div>
    </details>
  );
};
