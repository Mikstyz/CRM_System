import { EditableTitle } from '@/shared/ui/EditableTitle'
import { Discipline } from '../../entities/discipline/lib/types'

interface DisciplineItemProps {
  discipline: Discipline
  onDelete: () => void
}

export function DisciplineItem({ discipline, onDelete }: DisciplineItemProps) {
  const handleTitleSave = (newValue: string) => {
    console.log('Сохранённое название DisciplineItem:', newValue)
  }
  return (
    <div className="border rounded-lg p-2 flex justify-between items-center">
      <EditableTitle initialValue={discipline.title} onSave={handleTitleSave} />
      <div>{}</div>
      <button
        onClick={onDelete}
        className="bg-red-400 hover:bg-red-500 text-white font-bold py-1 px-2 rounded"
        title="Удалить дисциплину"
      >
        ✕
      </button>
    </div>
  )
}
