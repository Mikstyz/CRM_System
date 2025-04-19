import { useEffect, useState } from 'react'

interface NumberInputProps {
  value: number
  onChange: (value: number) => void
  min: number
  max: number
  classnames?: string
}

export function NumberInput({ value, onChange, min, max, classnames = '' }: NumberInputProps) {
  const [inputValue, setInputValue] = useState(String(value))

  // Синхронизируем локальное состояние с внешним значением
  useEffect(() => {
    setInputValue(String(value))
  }, [value])

  // Обработчик изменения ввода
  const handleChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    setInputValue(e.target.value)
  }

  // При потере фокуса валидируем и обновляем значение
  const handleBlur = () => {
    let numericValue = Number(inputValue.trim())

    // Если ввод пустой или некорректный, устанавливаем min
    if (!inputValue.trim() || Number.isNaN(numericValue)) {
      numericValue = min
    }

    // Ограничиваем числовое значение рамками
    numericValue = Math.max(min, Math.min(max, numericValue))

    // Обновляем значение только если оно изменилось
    if (numericValue !== value) {
      onChange(numericValue)
    }
    setInputValue(String(numericValue))
  }

  // Обработка нажатия клавиши Enter
  const handleKeyDown = (e: React.KeyboardEvent<HTMLInputElement>) => {
    if (e.key === 'Enter') {
      ;(e.target as HTMLInputElement).blur()
    }
  }

  return (
    <input
      type="number"
      value={inputValue}
      onChange={handleChange}
      onBlur={handleBlur}
      onKeyDown={handleKeyDown}
      min={min}
      max={max}
      className={`border rounded px-2 py-1 w-full ${classnames}`}
    />
  )
}
