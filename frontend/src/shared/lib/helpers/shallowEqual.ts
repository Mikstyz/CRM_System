export const shallowEqual = <T extends Record<string, unknown>>(a: T, b: T) =>
  Object.keys({ ...a, ...b }).every((k) => a[k] === b[k]);
