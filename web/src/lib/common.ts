export function formatPrice(price: number): string {
  if (isNaN(price) || price === 0) return "--";
  const before = Math.trunc(price / 100);
  const after = String(price % 100).padEnd(2, "0");
  return `R$ ${before},${after}`;
}

export const enum Page {
  Landing = "landing",
  Story = "story",
}

export function timeAgo(dateInput: string) {
  const past: Date = new Date(dateInput);
  const diffMs: number = Date.now() - past.getTime();

  const diffMinutes: number = Math.floor(diffMs / (1000 * 60));
  const diffHours: number = Math.floor(diffMs / (1000 * 60 * 60));
  const diffDays: number = Math.floor(diffMs / (1000 * 60 * 60 * 24));

  if (diffMinutes < 10) {
    return "just now";
  }

  if (diffMinutes < 60) {
    return "less than one hour ago";
  }

  if (diffHours < 24) {
    return diffHours === 1 ? "1 hour ago" : `${diffHours} hours ago`;
  }

  return diffDays === 1 ? "1 day ago" : `${diffDays} days ago`;
}
