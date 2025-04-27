import { FixedSizeList } from "react-window";
import { Discipline, Semester } from "@/entities/discipline/types";
import { DisciplineItem } from "@/widgets/DisciplineItem";
import { Id } from "@/shared/types";
import { useAppDispatch } from "@/shared/lib/hooks/redux.ts";
import { headerDeleteDiscipline } from "@/features/SemesterDisciplines/heanders/deleteDiscipline.ts";
import { headerAddDiscipline } from "@/features/SemesterDisciplines/heanders/addDiscipline.ts";

interface Props {
  semester: Semester;
  items: Discipline[];
  groupId: Id;
}

export const SemesterDisciplines = ({ semester, items, groupId }: Props) => {
  const dispatch = useAppDispatch();
  const itemData = items.map((d) => ({
    discipline: d,
    onDelete: () =>
      headerDeleteDiscipline({ dispatch, groupId, semester, discId: d.id }),
  }));

  return (
    <details className="border rounded-xl">
      <summary className="cursor-pointer select-none py-1 px-3 bg-slate-100 rounded-xl">
        {semester}‑й семестр
      </summary>

      <div className="p-3 space-y-2">
        <button
          onClick={() =>
            headerAddDiscipline({
              dispatch,
              groupId,
              semester,
              newTitle: "Новый предмет",
            })
          }
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
