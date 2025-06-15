import { ToastContainer } from "react-toastify";
import "react-toastify/dist/ReactToastify.css";

/** Единая точка инициализации Toastify */
export const WithToastify = () => (
  <ToastContainer
    position="top-right" // стандартная позиция
    autoClose={4000} // закроется сам через 4 с
    theme="colored" // цветная тема
    limit={3} // максимум 3 одновременных тоста
    newestOnTop
    pauseOnFocusLoss={false}
  />
);
