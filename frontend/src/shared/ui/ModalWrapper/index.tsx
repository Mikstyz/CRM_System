import { ReactNode, useEffect } from "react";
import { createPortal } from "react-dom";
import clsx from "classnames";

/**
 * ModalWrapper - Компонент-обёртка для модального окна.
 *
 * Этот компонент предоставляет возможность отображения модального окна,
 * которое адаптируется по ширине и высоте относительно содержимого.
 *
 * - Центрируется на экране.
 * - Ограничивает ширину и высоту для предотвращения переполнения.
 * - Поддерживает прокрутку при переполнении содержимого.
 *
 * Поддерживает закрытие по клику на фон.
 *
 * @component
 *
 * @param {Object} props - Свойства компонента.
 * @param {boolean} props.isOpen - Определяет, открыто ли модальное окно.
 * Если значение `false`, компонент не рендерится.
 * @param {function} props.onClose - Функция-обработчик, вызываемая при закрытии окна.
 * Вызывается при клике на фон или на кнопку закрытия.
 * @param {React.ReactNode} props.children - Контент, который будет отображён внутри модального окна.
 *
 * @example
 * Пример использования компонента ModalWrapper
 *
 * const App = () => {
 *   const [isModalOpen, setIsModalOpen] = useState(false);
 *
 *   return (
 *     <>
 *       <button onClick={() => setIsModalOpen(true)}>Открыть модальное окно</button>
 *       <ModalWrapper isOpen={isModalOpen} onClose={() => setIsModalOpen(false)}>
 *         <h1>Пример содержимого</h1>
 *       </ModalWrapper>
 *     </>
 *   );
 * };
 *
 */

interface ModalWrapperProps {
  isOpen: boolean;
  onClose(): void;
  children: ReactNode;
}

export function ModalWrapper({ isOpen, onClose, children }: ModalWrapperProps) {
  // ⛑️ SSR‑guard (если Wails/Next рендерят на сервере)
  if (typeof document === "undefined") return null;

  // гарантируем, что контейнер существует
  const modalRoot = document.getElementById("modal-root") ?? createModalRoot();

  // блокируем скролл и закрываем по Esc
  useEffect(() => {
    if (!isOpen) return;

    const onEsc = (e: KeyboardEvent) => e.key === "Escape" && onClose();
    document.body.style.overflow = "hidden";
    window.addEventListener("keydown", onEsc);

    return () => {
      document.body.style.overflow = "";
      window.removeEventListener("keydown", onEsc);
    };
  }, [isOpen, onClose]);

  if (!isOpen) return null;

  return createPortal(
    <div
      className="fixed inset-0 z-50 flex items-center justify-center bg-black/50"
      onClick={onClose}
    >
      <div
        className={clsx(
          "bg-white rounded-xl shadow-xl p-6",
          "max-h-[90vh] overflow-y-auto",
        )}
        onClick={(e) => e.stopPropagation()} /* внутри – не закрывать */
      >
        {children}
      </div>
    </div>,
    modalRoot,
  );
}

function createModalRoot() {
  const div = document.createElement("div");
  div.id = "modal-root";
  document.body.appendChild(div);
  return div;
}
