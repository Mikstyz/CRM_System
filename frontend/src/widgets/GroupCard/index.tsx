import { Group } from '../../entities/group/lib/types'
import { DisciplineItem } from '../DisciplineItem'
import { LinkDocument } from '@/shared/ui/LinkDocument'
import { PanelGroupCard } from '@/features/PanelGroupCard'
import { FixedSizeList, ListChildComponentProps } from 'react-window'
import { EditableTitle } from '@/shared/ui/EditableTitle'

interface GroupCardProps {
  group: Group
  onToggleExpand: () => void
  onAddDiscipline: () => void
  onDeleteGroup: () => void
  onDeleteDiscipline: (discId: string) => void
}

export function GroupCard({
  group,
  onToggleExpand,
  onAddDiscipline,
  onDeleteGroup,
  onDeleteDiscipline,
}: GroupCardProps) {
  const { name, isExpanded, disciplines } = group

  const handleTitleSave = (newValue: string) => {
    console.log('Сохранённое название GroupCard:', newValue)
  }

  const Row = ({ index, style }: ListChildComponentProps) => {
    const disc = disciplines[index]
    return (
      <div style={style}>
        <DisciplineItem key={disc.id} discipline={disc} onDelete={() => onDeleteDiscipline(disc.id)} />
      </div>
    )
  }
  return (
    <div className="border-2 rounded-xl p-4">
      <div className="flex items-center justify-between mb-2">
        <div className="text-lg font-semibold">
          Группа:{' '}
          <span className="ml-1">
            <EditableTitle initialValue={name} onSave={handleTitleSave} className="w-min" />
          </span>
        </div>
        <PanelGroupCard onToggleExpand={onToggleExpand} onDeleteGroup={onDeleteGroup} isExpanded={isExpanded} />
      </div>

      {isExpanded && (
        <div className="border-t pt-4 mt-2">
          <div className="flex items-center justify-between mb-2">
            <button
              onClick={onAddDiscipline}
              className="bg-green-400 hover:bg-green-500 text-white font-bold py-1 px-4 rounded"
            >
              +
            </button>
            <LinkDocument href="#">Открыть xlsx дисциплин</LinkDocument>
          </div>

          {/* Список дисциплин */}
          <FixedSizeList height={300} itemCount={disciplines.length} itemSize={55} width="100%">
            {Row}
          </FixedSizeList>
        </div>
      )}
    </div>
  )
}
