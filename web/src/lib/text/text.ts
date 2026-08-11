import translations from "../../assets/texts.json";

export function getTextDefault(k: string, d: string): string {
  return translations[k as keyof typeof translations] || d;
}

export function getText(k: string): string {
  return getTextDefault(k, "<unknown text>");
}
