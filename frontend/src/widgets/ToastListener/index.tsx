import { useEffect, useRef } from "react";
import { toast, Id, ToastOptions } from "react-toastify";
import { shallowEqual } from "react-redux";
import { useAppDispatch, useAppSelector } from "@/shared/lib/hooks/redux";
import {
  clearErrors as clearGroups,
  clearMessage as clearGroupsMsg,
} from "@/entities/group/store";
import { blankClearErrors, blankClearMessage } from "@/entities/blank/store";

const PENDING_LABEL = "Загрузка...";

export const ToastListener = () => {
  const dispatch = useAppDispatch();

  const shown = useRef<Record<string, Id>>({});

  const { grpErr, grpMsg, blkErr, blkMsg } = useAppSelector(
    (s) => ({
      grpErr: s.groups.error,
      grpMsg: s.groups.loading?.message,
      blkErr: s.blank.error,
      blkMsg: s.blank.loading?.message,
    }),
    shallowEqual,
  );

  /** Утилита отображения */
  const pushToast = (
    key: string,
    content: string,
    opts: ToastOptions,
    clearAction: () => void,
  ) => {
    if (shown.current[key] && toast.isActive(shown.current[key]!)) return;

    shown.current[key] = toast(content, {
      ...opts,
      toastId: key,
      onClose: () => {
        clearAction();
        delete shown.current[key];
      },
    });
  };

  /* === Ошибки ============================================================ */
  useEffect(() => {
    if (grpErr) {
      pushToast(
        `err-grp-${grpErr}`,
        grpErr,
        { type: "error", autoClose: 2000 },
        () => dispatch(clearGroups()),
      );
    }
  }, [grpErr, dispatch]);

  useEffect(() => {
    if (blkErr) {
      pushToast(
        `err-blk-${blkErr}`,
        blkErr,
        { type: "error", autoClose: 2000 },
        () => dispatch(blankClearErrors()),
      );
    }
  }, [blkErr, dispatch]);

  /* === Успешные сообщения =============================================== */
  useEffect(() => {
    if (grpMsg && grpMsg !== PENDING_LABEL) {
      pushToast(
        `msg-grp-${grpMsg}`,
        grpMsg,
        { type: "success", autoClose: 2000 },
        () => dispatch(clearGroupsMsg()),
      );
    }
  }, [grpMsg, dispatch]);

  useEffect(() => {
    if (blkMsg && blkMsg !== PENDING_LABEL) {
      pushToast(
        `msg-blk-${blkMsg}`,
        blkMsg,
        { type: "success", autoClose: 2000 },
        () => dispatch(blankClearMessage()),
      );
    }
  }, [blkMsg, dispatch]);

  return null;
};
