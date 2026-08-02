export function formatPrice(price: number): string {
  if (isNaN(price) || price === 0) return "--";
  const before = Math.trunc(price / 100);
  const after = String(price % 100).padEnd(2, "0");
  return `R$ ${before},${after}`;
}
