import clsx from "classnames";

interface Props {
  title: string;
  inputProps: React.InputHTMLAttributes<HTMLInputElement>;
  error?: string;
}

export function InputFilterGroup({title, inputProps, error}: Props) {
  return (
    <label className="flex flex-col w-[48%] min-w-[120px]">
      <span className="mb-1">{title}</span>
      <input {...inputProps}
             className={clsx(
        "border rounded px-3 py-2 w-full focus:outline-none focus:ring-1",
               error ? "border-red-500 focus:border-green-500" : "border-gray-300",
      )}/>
      {error && <span className="text-red-500 text-xs mt-1">{error}</span>}
    </label>
  );
}
