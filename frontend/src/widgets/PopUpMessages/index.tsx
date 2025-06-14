import { useAppDispatch, useAppSelector } from "@/shared/lib/hooks/redux.ts";
import { useEffect, useState } from "react";
import { clearErrors as clearGroups } from "@/entities/group/store";
import { blankClearErrors } from "@/entities/blank/store";

export function PopUpMessages() {
  const dispatch = useAppDispatch();
  const grpErr = useAppSelector((s) => s.groups.error);
  const blankErr = useAppSelector((s) => s.blank.error);
  const [visible, setVisible] = useState<string | null>(null);

  useEffect(() => {
    const err = grpErr || blankErr;
    if (!err) return;

    setVisible(err);

    const t = setTimeout(() => {
      setVisible(null);
      if (grpErr) dispatch(clearGroups());
      if (blankErr) dispatch(blankClearErrors());
    }, 4_000);

    return () => clearTimeout(t);
  }, [grpErr, blankErr, dispatch]);

  if (!visible) return null;

  return (
    <div className="fixed top-4 right-4 bg-red-600 text-white px-4 py-2 rounded shadow-lg z-100">
      {visible}
    </div>
  );
}
