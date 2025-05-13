import { Component, ReactNode } from "react";

export class ModalErrorBoundary extends Component<
  { children: ReactNode },
  { hasError: boolean }
> {
  state = { hasError: false };

  static getDerivedStateFromError() {
    return { hasError: true };
  }

  componentDidCatch(err: unknown, info: unknown) {
    console.error("Modal crashed:", err, info);
  }

  render() {
    return this.state.hasError ? (
      <p className="p-4 text-red-600">Что‑то пошло не так…</p>
    ) : (
      this.props.children
    );
  }
}
