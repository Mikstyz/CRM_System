import { Loading } from "@/shared/types/store.type.ts";

interface BaseState {
  loading: Loading;
  error?: string;
}

export const handlePending = <S extends BaseState>(state: S): void => {
  state.loading = {
    status: "pending",
    message: "Загрузка...",
  };
  state.error = undefined;
};

export const handleRejected = <S extends BaseState>(
  state: S,
  action: { payload?: string },
): void => {
  state.loading = {
    status: "failed",
    message: "",
  };
  state.error = action.payload || "Произошла ошибка";
};
