import { ButtonPanelGroupCard } from "./ui/ButtonPanelGroupCard";

export function PanelGroupCard({
  onToggleExpand,
  onDeleteGroup,
  onDuplicateGroup,
  onOpenBlank,
  isExpanded,
}: {
  onToggleExpand: () => void;
  onDeleteGroup: () => void;
  onDuplicateGroup: () => void;
  onOpenBlank: () => void;
  isExpanded: boolean;
}) {
  return (
    <div className="flex items-center space-x-2">
      <ButtonPanelGroupCard
        onClick={onToggleExpand}
        title="Свернуть/Развернуть"
        className="bg-green-400 hover:bg-green-500"
      >
        {isExpanded ? "▲" : "▼"}
      </ButtonPanelGroupCard>
      <ButtonPanelGroupCard
        onClick={onOpenBlank}
        title="Открыть бланк"
        className="bg-yellow-400 hover:bg-yellow-500"
      >
        📁
      </ButtonPanelGroupCard>
      <ButtonPanelGroupCard
        onClick={onDuplicateGroup}
        title="Дублировать группу"
        className="bg-orange-400 hover:bg-orange-500"
      >
        📄
      </ButtonPanelGroupCard>
      <ButtonPanelGroupCard
        onClick={onDeleteGroup}
        title="Удалить группу"
        className="bg-red-400 hover:bg-red-500"
      >
        ✕
      </ButtonPanelGroupCard>
    </div>
  );
}
