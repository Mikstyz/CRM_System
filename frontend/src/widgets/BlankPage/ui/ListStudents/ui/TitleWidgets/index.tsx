import { EditableTitle } from "src/shared/ui/EditableTitle";
import { handleTitleGroupSave } from "@/shared/lib/headers/titleGroupSave.ts";
import { useState } from "react";
import { useAppDispatch } from "@/shared/lib/hooks/redux.ts";
import { Group } from "@/entities/group/types";

interface TitleWidgetsProps {
  group: Group;
}

export function TitleWidgets({ group }: TitleWidgetsProps) {
  const dispatch = useAppDispatch();
  const [err, setErr] = useState<string | undefined>(undefined);

  return (
    <div className="flex items-center justify-between">
      <h1 className="text-2xl font-bold">
        Группа:{" "}
        <EditableTitle
          value={group.name}
          onSave={async (value) => {
            const ok = await handleTitleGroupSave({
              dispatch,
              group,
              value,
            });
            setErr(ok ? undefined : "Неверный формат имени");
          }}
          error={err}
        />
      </h1>
    </div>
  );
}
