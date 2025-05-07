import React from "react";
import ReactDOM from "react-dom/client";
import { Provider } from "react-redux";
import "./styles/index.css";
import { store } from "@/app/store";
import { PagesListGroup } from "@/pages/PagesListGroup";
import { ConfirmProvider } from "@/shared/ui/ConfirmDialog";

ReactDOM.createRoot(document.getElementById("root") as HTMLElement).render(
  <React.StrictMode>
    <Provider store={store}>
      {/* для отображения Confirm перед действием */}
      <ConfirmProvider>
        <PagesListGroup />
      </ConfirmProvider>
    </Provider>
  </React.StrictMode>,
);
