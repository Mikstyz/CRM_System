import React, {
  createContext,
  useCallback,
  useContext,
  useEffect,
  useRef,
  useState,
  ReactNode,
} from "react";
import { createPortal } from "react-dom";
import { AnimatePresence, motion } from "framer-motion";
import classNames from "classnames";

/**
 * ---------------------------------------------------------------------------
 * ConfirmProvider / useConfirm / <ConfirmDialog />
 * ---------------------------------------------------------------------------
 * A promise‑based replacement for the native `window.confirm()` using
 * React19, TailwindCSSv4 and Framer Motion.  It fulfils several goals:
 * –Focus trap + `aria-modal` ==> the page becomes inaccessible until the user
 *   responds (meets WCAG2.2 «modal dialog» recommendations).
 * –Promise interface → the calling code reads naturally with `await confirm()`.
 * –Provider/Hook pattern keeps Feature‑Sliced architecture clean: the modal
 *   lives in /shared/ui, while features/pages just import `useConfirm()`.
 * –SOLID: `ConfirmDialog` (Single‑responsibility), `useConfirm` (I, D)
 * –Typed API with sensible defaults & full customisation.
 *
 * Usage ---------------------------------------------------------------------
 * 1.  Wrap **once** at the top of your tree (e.g. `src/app/providers`):
 *
 *     <ConfirmProvider>
 *       <App />
 *     </ConfirmProvider>
 *
 * 2.  Inside any component / handler:
 *
 *     const confirm = useConfirm();
 *     const ok = await confirm({
 *       title: "Удалить группу?",
 *       description: "Это действие нельзя отменить.",
 *     });
 *     if (ok) dispatch(deleteGroupsThunks(group.id));
 *
 * 3.  That’s it – no prop‑drilling, no Redux wiring.
 * ---------------------------------------------------------------------------
 */

/* ------------------------------------------------------------------------- *
 *                         Types & Context
 * ------------------------------------------------------------------------- */
export interface ConfirmOptions {
  /** Heading of the dialog */
  title?: string;
  /** Optional description paragraph */
  description?: string;
  /** «Yes» button label */
  confirmText?: string;
  /** «No» button label */
  cancelText?: string;
}

export type ConfirmFn = (options?: ConfirmOptions) => Promise<boolean>;

interface ConfirmState {
  visible: boolean;
  options: ConfirmOptions;
  resolver?: (value: boolean) => void;
}

const ConfirmContext = createContext<ConfirmFn | null>(null);

/* ------------------------------------------------------------------------- *
 *                           Provider
 * ------------------------------------------------------------------------- */
export const ConfirmProvider = ({ children }: { children: ReactNode }) => {
  const [state, setState] = useState<ConfirmState>({
    visible: false,
    options: {},
  });

  /**
   * Opens the dialog and returns a promise resolved with the user's choice.
   */
  const confirm = useCallback<ConfirmFn>(
    (options = {}) =>
      new Promise<boolean>((resolve) => {
        setState({ visible: true, options, resolver: resolve });
      }),
    [],
  );

  const handleClose = (result: boolean) => {
    state.resolver?.(result);
    setState((prev) => ({ ...prev, visible: false, resolver: undefined }));
  };

  return (
    <ConfirmContext.Provider value={confirm}>
      {children}
      {/* Dialog is rendered here so it exists exactly once */}
      <ConfirmDialog
        isOpen={state.visible}
        options={state.options}
        onClose={handleClose}
      />
    </ConfirmContext.Provider>
  );
};

/**
 * Hook to obtain the confirm function.
 * Throws if used outside <ConfirmProvider> – helps avoid silent errors.
 */
export const useConfirm = (): ConfirmFn => {
  const ctx = useContext(ConfirmContext);
  if (!ctx) throw new Error("useConfirm must be used within ConfirmProvider");
  return ctx;
};

/* ------------------------------------------------------------------------- *
 *                         Dialog component (private)
 * ------------------------------------------------------------------------- */
interface ConfirmDialogProps {
  isOpen: boolean;
  options: ConfirmOptions;
  onClose: (result: boolean) => void;
}

const ConfirmDialog = ({ isOpen, options, onClose }: ConfirmDialogProps) => {
  const dialogRef = useRef<HTMLDivElement>(null);

  /* ---------------- Focus trap & inert root ---------------- */
  useEffect(() => {
    if (!isOpen) return;

    // save previously focused element to restore later
    const previouslyFocused = document.activeElement as HTMLElement | null;

    // Mark the rest of the app as inert so screen‑readers / tabbing ignore it
    const appRoot = document.getElementById("root");
    appRoot?.setAttribute("inert", "");

    // Focus first button in the dialog
    setTimeout(() => dialogRef.current?.querySelector("button")?.focus());

    return () => {
      // Restore
      appRoot?.removeAttribute("inert");
      previouslyFocused?.focus();
    };
  }, [isOpen]);

  if (!isOpen) return null;

  return createPortal(
    <AnimatePresence>
      {/*     Backdrop     */}
      <motion.div
        className="fixed inset-0 z-50 flex items-center justify-center"
        initial={{ opacity: 0 }}
        animate={{ opacity: 1 }}
        exit={{ opacity: 0 }}
      >
        <div className="absolute inset-0 bg-black/50 backdrop-blur-sm" />

        {/* Dialog */}
        <motion.div
          role="dialog"
          aria-modal="true"
          aria-labelledby="confirm-title"
          aria-describedby="confirm-description"
          ref={dialogRef}
          tabIndex={-1}
          className={classNames(
            "relative z-10 w-full max-w-md rounded-2xl bg-white p-6 shadow-xl",
            "flex flex-col",
          )}
          initial={{ scale: 0.9, opacity: 0 }}
          animate={{ scale: 1, opacity: 1 }}
          exit={{ scale: 0.9, opacity: 0 }}
        >
          <h2
            id="confirm-title"
            className="text-lg font-semibold text-gray-900"
          >
            {options.title ?? "Confirm"}
          </h2>
          {options.description && (
            <p id="confirm-description" className="mt-2 text-sm text-gray-600">
              {options.description}
            </p>
          )}

          <div className="mt-6 flex justify-end space-x-3">
            <button
              onClick={() => onClose(false)}
              className="rounded-lg border border-gray-300 px-4 py-2 text-sm font-medium text-gray-700 hover:bg-gray-100 focus:outline-none focus:ring-2 focus:ring-gray-400"
            >
              {options.cancelText ?? "No"}
            </button>
            <button
              onClick={() => onClose(true)}
              className="rounded-lg bg-blue-600 px-4 py-2 text-sm font-medium text-white shadow hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-blue-400"
            >
              {options.confirmText ?? "Yes"}
            </button>
          </div>
        </motion.div>
      </motion.div>
    </AnimatePresence>,
    document.body,
  );
};
