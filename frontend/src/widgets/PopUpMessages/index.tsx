import { useAppDispatch, useAppSelector } from "@/shared/lib/hooks/redux.ts";
import { useEffect, useState } from "react";
import {
  clearErrors as clearGroups,
  clearMessage as clearGroupsMessage,
} from "@/entities/group/store";
import { blankClearErrors, blankClearMessage } from "@/entities/blank/store";
export function PopUpMessages() {
  const dispatch = useAppDispatch();
  const grpErr = useAppSelector((s) => s.groups.error);
  const blankErr = useAppSelector((s) => s.blank.error);
  const grpMsg = useAppSelector((s) => s.groups.loading.message);
  const blankMsg = useAppSelector((s) => s.blank.loading.message);
  const [visible, setVisible] = useState<{
    text: string;
    type: "error" | "success";
  } | null>(null);

  useEffect(() => {
    const err = grpErr || blankErr;
    if (!err) return;

    setVisible({ text: err, type: "error" });

    const t = setTimeout(() => {
      setVisible(null);
      if (grpErr) dispatch(clearGroups());
      if (blankErr) dispatch(blankClearErrors());
    }, 4_000);

    return () => clearTimeout(t);
  }, [grpErr, blankErr, dispatch]);

  useEffect(() => {
    const msg = grpMsg || blankMsg;
    if (!msg || msg === "Загрузка...") return;

    setVisible({ text: msg, type: "success" });

    const t = setTimeout(() => {
      setVisible(null);
      if (grpMsg) dispatch(clearGroupsMessage());
      if (blankMsg) dispatch(blankClearMessage());
    }, 2_000);

    return () => clearTimeout(t);
  }, [grpMsg, blankMsg, dispatch]);

  if (!visible) return null;

  const bg = visible.type === "error" ? "bg-red-600" : "bg-green-600";

  return (
    <div
      className={`fixed top-4 right-4 text-white px-4 py-2 rounded shadow-lg z-100 ${bg}`}
    >
      {visible.text}
    </div>
  );
}
