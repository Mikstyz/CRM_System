export { handlePending, handleRejected } from './StoreHandlers'

/**
 * Удаляет карточку из массива по переданному идентификатору
 * @param items - массив объектов-карточек
 * @param id - идентификатор удаляемой карточки
 * @returns новый массив без удаленного элемента
 */
export function removeCard<T extends { id: string | number }>(items: T[], id: string | number): T[] {
  return items.filter((item) => item.id !== id)
}
