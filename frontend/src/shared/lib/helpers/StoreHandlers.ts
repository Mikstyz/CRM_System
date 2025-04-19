/**
 * Универсальные обработчики для pending/rejected.
 */
export const handlePending = (state: any) => {
  state.loading = true
  state.error = null
}

export const handleRejected = (state: any, action: any) => {
  state.loading = false
  state.error = action.payload || 'Произошла ошибка'
}
