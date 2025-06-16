import React, {
  useState,
  useRef,
  useEffect,
  ChangeEvent,
  KeyboardEvent,
} from "react";
import clsx from "classnames";

interface EditableTitleProps {
  value?: string;
  onSave?: (newValue: string) => void;
  className?: string;
  error?: string;
}

export function EditableTitle({
  value = "",
  onSave,
  className = "",
  error,
}: EditableTitleProps) {
  const [draft, setDraft] = useState(value);

  const [edit, setEdit] = useState(false);
  const inputRef = useRef<HTMLInputElement>(null);

  // auto-focus при переходе в режим редактирования
  useEffect(() => {
    if (edit && inputRef.current) {
      inputRef.current.focus();
      inputRef.current.select();
    }
  }, [edit]);

  const commit = () => {
    setEdit(false);
    if (draft != null) {
      onSave?.(draft.trim());
    }
  };

  const cancel = () => {
    setDraft(value);
    setEdit(false);
  };

  const onKeyDown = (e: KeyboardEvent) => {
    if (e.key === "Enter") commit();
    if (e.key === "Escape") cancel();
  };

  return (
    <div
      className={clsx("inline-flex items-center group", className)}
      onClick={() => !edit && setEdit(true)}
      role="button"
      tabIndex={0}
      aria-label="Нажмите, чтобы редактировать заголовок"
      onKeyDown={(e) => e.key === "Enter" && setEdit(true)}
    >
      {!edit && (
        <>
          <span className="editable-hover">{draft || "—"}</span>
          <svg
            aria-hidden="true"
            className="ml-1 h-4 w-4 opacity-0 group-hover:opacity-60 transition-opacity pointer-events-none"
            viewBox="0 0 20 20"
          >
            <path d="M2 14.5l2.6.6 9-9-2.6-2.6-9 9z" />
          </svg>
        </>
      )}

      {edit && (
        <input
          ref={inputRef}
          type="text"
          value={draft}
          onChange={(e: ChangeEvent<HTMLInputElement>) =>
            setDraft(e.target.value)
          }
          onBlur={commit}
          onKeyDown={onKeyDown}
          className={clsx(
            "border-b focus:outline-none bg-transparent caret-current",
            error ? "border-red-500" : "border-current",
          )}
          role="textbox"
          aria-invalid={!!error}
          aria-describedby={error ? "title-error" : undefined}
        />
      )}

      {error && (
        <p
          id="title-error"
          className="ml-2 text-xs text-red-500"
          aria-live="polite"
        >
          {error}
        </p>
      )}
    </div>
  );
}

//import React, { ForwardedRef, forwardRef, useMemo, useState } from "react";
// import {
//   Box,
//   Typography,
//   TextField,
//   Tooltip,
//   IconButton,
//   SxProps,
//   Theme,
//   TypographyProps,
//   BoxProps,
// } from "@mui/material";
// import EditIcon from "@mui/icons-material/Edit";
//
// export interface InlineEditableProps extends BoxProps {
//   value: string;
//   onSave?: (v: string) => void;
//   variant?: TypographyProps["variant"];
//   error?: string;
// }
//
// export const EditableTitle = forwardRef<HTMLDivElement, InlineEditableProps>(
//   function EditableTitle(
//     { value, onSave, variant = "h6", error, sx: sxProp, ...boxProps },
//     ref: ForwardedRef<HTMLDivElement>,
//   ) {
//     const [editing, setEditing] = useState(false);
//     const [draft, setDraft] = useState(value);
//
//     const base: SxProps<Theme> = {
//       display: "inline-flex",
//       alignItems: "center",
//       "&:hover .affordance": {
//         textDecoration: "underline dotted",
//         opacity: 1,
//       },
//     };
//
//     const mergedSx: SxProps<Theme> = useMemo(() => {
//       const extra =
//         sxProp == null ? [] : Array.isArray(sxProp) ? sxProp : [sxProp];
//       return [base, ...extra];
//     }, [sxProp]);
//
//     const commit = () => {
//       setEditing(false);
//       if (draft.trim() !== value) onSave?.(draft.trim());
//     };
//
//     return (
//       <Box ref={ref} sx={mergedSx} {...boxProps}>
//         {editing ? (
//           <TextField
//             variant="standard"
//             value={draft}
//             onChange={(e) => setDraft(e.target.value)}
//             onBlur={commit}
//             onKeyDown={(e) => e.key === "Enter" && commit()}
//             autoFocus
//             error={!!error}
//             helperText={error}
//             sx={{ minWidth: 80 }}
//           />
//         ) : (
//           <>
//             <Typography
//               variant={variant}
//               component="span"
//               className="affordance"
//               sx={{ cursor: "pointer", transition: "text-decoration 120ms" }}
//               tabIndex={0}
//               onClick={() => setEditing(true)}
//               onKeyDown={(e) =>
//                 (e.key === "Enter" || e.key === " ") && setEditing(true)
//               }
//             >
//               {value || "—"}
//             </Typography>
//             <Tooltip title="Редактировать" arrow>
//               <IconButton
//                 size="small"
//                 className="affordance"
//                 sx={{ ml: 0.5, opacity: 0, transition: "opacity 120ms" }}
//                 aria-label="Редактировать"
//               >
//                 <EditIcon fontSize="inherit" />
//               </IconButton>
//             </Tooltip>
//           </>
//         )}
//       </Box>
//     );
//   },
// );
//
// EditableTitle.displayName = "EditableTitle";
