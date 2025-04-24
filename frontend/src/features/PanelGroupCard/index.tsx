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
        {isExpanded ? "▼" : "▲"}
      </ButtonPanelGroupCard>
      <ButtonPanelGroupCard
        onClick={onOpenBlank}
        title="Открыть бланк"
        className="bg-yellow-500 hover:bg-yellow-600"
      >
        <svg
          fill="#000000"
          height="75%"
          width="75%"
          viewBox="0 0 317.001 317.001"
        >
          <g
            fill="currentColor"
            stroke="currentColor"
            strokeWidth="10"
            transform="translate(2 2)"
          >
            <path
              d="M270.825,70.55L212.17,3.66C210.13,1.334,207.187,0,204.093,0H55.941C49.076,0,43.51,5.566,43.51,12.431V304.57
	c0,6.866,5.566,12.431,12.431,12.431h205.118c6.866,0,12.432-5.566,12.432-12.432V77.633
	C273.491,75.027,272.544,72.51,270.825,70.55z M55.941,305.073V12.432H199.94v63.601c0,3.431,2.78,6.216,6.216,6.216h54.903
	l0.006,222.824H55.941z"
            />
          </g>
        </svg>
      </ButtonPanelGroupCard>
      <ButtonPanelGroupCard
        onClick={onDuplicateGroup}
        title="Дублировать группу"
        className="bg-orange-500 hover:bg-orange-600"
      >
        <svg height="75%" width="75%" viewBox="0 0 21 21">
          <g
            fill="none"
            stroke="currentColor"
            strokeWidth="2"
            transform="translate(2 2)"
          >
            <path d="m12.5 14.5v-8c0-1.1045695-.8954305-2-2-2h-8c-1.1045695 0-2 .8954305-2 2v8c0 1.1045695.8954305 2 2 2h8c1.1045695 0 2-.8954305 2-2z" />
            <path d="m12.5 12.5h2c1.1045695 0 2-.8954305 2-2v-7.99654173c0-1.1045695-.8954305-2-2-2-.0011518 0-.0023035 0-.0034553 0l-7.99999998.01382415c-1.10321873.00190597-1.99654472.89677664-1.99654472 1.99999701v1.98272057" />
            <path d="m6.5 7.5v6" />
            <path d="m6.5 7.5v6" transform="matrix(0 1 -1 0 17 4)" />
          </g>
        </svg>
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
