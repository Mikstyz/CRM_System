interface BaseState {
  loading: boolean;
  error: string | null;
}

export const handlePending = <S extends BaseState>(state: S): void => {
  state.loading = true;
  state.error = null;
};

export const handleRejected = <S extends BaseState>(
  state: S,
  action: { payload?: string },
): void => {
  state.loading = false;
  state.error = action.payload || "Произошла ошибка";
};
