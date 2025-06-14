import styles from "./index.module.css";
import classNames from "classnames";

interface UsageInformationProps {
  className?: string;
}
export function UsageInformation({ className }: UsageInformationProps) {
  return (
    <aside
      className={classNames(
        className,
        styles.root,
        "bg-gray-800 text-gray-200 rounded-xl shadow-lg flex flex-col",
      )}
    >
      <header className="flex items-center mb-4">
        <h2 className="text-lg font-semibold">Инструкция</h2>
      </header>

      <ol className="list-decimal list-inside space-y-2 text-sm leading-relaxed flex-1 overflow-auto">
        <li>Выберите курс и специальность.</li>
        <li>Укажите число выпускников.</li>
        <li>Номер группы вводится цифровым спиннером.</li>
        <li>Нажмите «Добавить» для сохранения.</li>
      </ol>
    </aside>
  );
}
