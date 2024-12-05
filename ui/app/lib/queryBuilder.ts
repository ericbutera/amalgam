import { Pagination } from "@/app/types/pagination";

export default function queryString(
  pagination: Pagination,
  overrides: Partial<Pagination> = {},
): string {
  const q = new URLSearchParams();
  const p = { ...pagination, ...overrides };

  if (p.previous) {
    q.append("previous", p.previous);
  }

  if (p.next) {
    q.append("next", p.next);
  }

  if (p.limit !== undefined) {
    q.append("limit", String(p.limit));
  }

  return q.toString();
}
