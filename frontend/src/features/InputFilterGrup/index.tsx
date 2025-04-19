interface InputFilterGrupProps {
  onChange: (event: React.ChangeEvent<HTMLInputElement>) => void
  value: string
  title: string
}

export function InputFilterGrup({ onChange, value, title }: InputFilterGrupProps) {
  return (
    <div className="w-[40%] min-w-50">
      <label className="block mb-1">{title}</label>
      <input type="text" value={value} onChange={onChange} className="w-full mt-1 p-1 border rounded" />
    </div>
  )
}
