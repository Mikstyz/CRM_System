import React from "react";
import ReactDOM from "react-dom/client";
import { Provider } from "react-redux";
import "./styles/index.css";
import "@/app/styles/normolaz.css";
import { store } from "@/app/store";
import { PagesListGroup } from "@/pages/PagesListGroup";
import { ConfirmProvider } from "@/shared/ui/ConfirmDialog";
import { WithToastify } from "@/app/providers/with-toastify.tsx";
import { ToastListener } from "@/widgets/ToastListener";

ReactDOM.createRoot(document.getElementById("root") as HTMLElement).render(
  <React.StrictMode>
    <Provider store={store}>
      {/* для отображения Confirm перед действием */}
      <ConfirmProvider>
        <PagesListGroup />
      </ConfirmProvider>
      {/* Слушатель ошибок/сообщений */}
      <ToastListener />

      {/* Контейнер уведомлений */}
      <WithToastify />
    </Provider>
  </React.StrictMode>,
);
