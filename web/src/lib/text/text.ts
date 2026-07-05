import translations from "../../assets/texts.json";

export function getText(k: string): string {
  return translations[k as keyof typeof translations] || "<unknown text>";
}
