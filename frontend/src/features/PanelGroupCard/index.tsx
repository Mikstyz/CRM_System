import { useNavigate } from 'react-router-dom'
import { ButtonPanelGroupCard } from './ui/ButtonPanelGroupCard'

export function PanelGroupCard({
  onToggleExpand,
  onDeleteGroup,
  isExpanded,
}: {
  onToggleExpand: () => void
  onDeleteGroup: () => void
  isExpanded: boolean
}) {
  const navigation = useNavigate()

  return (
    <div className="flex items-center space-x-2">
      <ButtonPanelGroupCard
        onClick={onToggleExpand}
        title="Свернуть/Развернуть"
        className="bg-green-400 hover:bg-green-500"
      >
        {isExpanded ? '▲' : '▼'}
      </ButtonPanelGroupCard>
      <ButtonPanelGroupCard
        onClick={() => navigation('/blank/')}
        title="Открыть бланк"
        className="bg-yellow-400 hover:bg-yellow-500"
      >
        📁
      </ButtonPanelGroupCard>
      <ButtonPanelGroupCard
        onClick={() => {
          console.log('Дублируем группу')
        }}
        title="Дублировать группу"
        className="bg-orange-400 hover:bg-orange-500"
      >
        📄
      </ButtonPanelGroupCard>
      <ButtonPanelGroupCard onClick={onDeleteGroup} title="Удалить группу" className="bg-red-400 hover:bg-red-500">
        ✕
      </ButtonPanelGroupCard>
    </div>
  )
}
