import React from "react";
import ReactDOM from "react-dom/client";
import { Provider } from "react-redux";
import "./styles/index.css";
import { store } from "@/app/store";
import { PagesListGroup } from "@/pages/PagesListGroup";

ReactDOM.createRoot(document.getElementById("root") as HTMLElement).render(
  <React.StrictMode>
    <Provider store={store}>
      <PagesListGroup />
    </Provider>
  </React.StrictMode>,
);
