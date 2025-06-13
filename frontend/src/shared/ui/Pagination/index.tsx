import clsx from "classnames";

interface Props {
  page: number;
  pages: number;
  onChange: (p: number) => void;
}

export function Pagination({ page, pages, onChange }: Props) {
  if (pages <= 1) return null;

  const go = (p: number) => {
    if (p >= 1 && p <= pages && p !== page) onChange(p);
  };

  return (
    <nav className="flex items-center justify-center gap-2 select-none">
      <button
        onClick={() => go(page - 1)}
        className={clsx("px-2", page === 1 && "opacity-40 cursor-default")}
      >
        ‹
      </button>

      {Array.from({ length: pages }, (_, i) => i + 1).map((p) => (
        <button
          key={p}
          onClick={() => go(p)}
          className={clsx(
            "px-2 py-1 rounded",
            p === page
              ? "bg-blue-600 text-white"
              : "hover:bg-gray-100 text-gray-700",
          )}
        >
          {p}
        </button>
      ))}

      <button
        onClick={() => go(page + 1)}
        className={clsx("px-2", page === pages && "opacity-40 cursor-default")}
      >
        ›
      </button>
    </nav>
  );
}
