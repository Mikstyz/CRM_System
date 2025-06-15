export function isoToRu(iso?: string): string {
  if (!iso) return "";
  const parts = iso.split("-");
  if (parts.length !== 3) return iso;
  const [y, m, d] = parts;
  return `${d}.${m}.${y}`;
}

export function ruToIso(ru?: string): string {
  if (!ru) return "";
  if (ru.includes("-")) return ru; // already ISO
  const parts = ru.split(".");
  if (parts.length !== 3) return ru;
  const [d, m, y] = parts;
  return `${y}-${m}-${d}`;
}
