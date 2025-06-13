import React, { useState, ChangeEvent } from "react";

interface EditableTitleProps {
  initialValue?: string;
  onSave?: (newValue: string) => void;
  className?: string;
  error?: string;
}

/**
 * Компонент, который выглядит как заголовок, но позволяет менять текст.
 * @param initialValue Стартовое значение заголовка.
 * @param onSave Колбэк, вызывается при потере фокуса или на Enter.
 * @param className
 * @param error
 */
export function EditableTitle({
  initialValue = "",
  onSave,
  className = "",
  error,
}: EditableTitleProps) {
  const [value, setValue] = useState(initialValue);

  const handleChange = (event: ChangeEvent<HTMLInputElement>) => {
    setValue(event.target.value);
  };

  const handleBlurOrEnter = () => {
    if (onSave) {
      onSave(value);
    }
  };

  const handleKeyDown = (event: React.KeyboardEvent<HTMLInputElement>) => {
    if (event.key === "Enter") {
      // Снимаем фокус, чтобы сработал onBlur
      (event.target as HTMLInputElement).blur();
    }
  };

  return (
    <>
      <input
        type="text"
        className={`font-bold outline-none focus:ring-0 border-b
          ${error ? "border-red-500" : "border-transparent"} ${className}`}
        value={value}
        onChange={handleChange}
        onBlur={handleBlurOrEnter}
        onKeyDown={handleKeyDown}
      />
      {error && <p className="text-xs text-red-500">{error}</p>}
    </>
  );
}
