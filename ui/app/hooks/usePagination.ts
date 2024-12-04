import { useSearchParams } from "next/navigation";
import { Pagination } from "@/app/types/pagination";

export function usePagination(): Pagination {
  const searchParams = useSearchParams();
  return {
    previous: searchParams.get("previous") || undefined,
    next: searchParams.get("next") || undefined,
    limit: searchParams.get("limit")
      ? Number(searchParams.get("limit"))
      : undefined,
  };
}
