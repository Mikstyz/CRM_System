interface Props {
  title: string;
  inputProps: React.InputHTMLAttributes<HTMLInputElement>;
  error?: string;
}

export function InputFilterGroup({ title, inputProps, error }: Props) {
  return (
    <label className="flex flex-col w-[48%] min-w-[120px]">
      <span className="mb-1">{title}</span>
      <input {...inputProps} />
      {error && <span className="text-red-500 text-xs mt-1">{error}</span>}
    </label>
  );
}
