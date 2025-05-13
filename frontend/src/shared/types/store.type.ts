export type Loading = {
  status: "idle" | "pending" | "succeeded" | "failed";
  message?: string;
};
