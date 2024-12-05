import useSWR from "swr";
import { getGraph } from "../lib/fetch";
import { Pagination } from "../types/pagination";
import queryString from "../lib/queryBuilder";

export default function useArticles(feedId: string, pagination: Pagination) {
  const fetcher = async () => {
    return await getGraph().ListArticles({
      feedId,
      previous: pagination.previous,
      next: pagination.next,
      limit: pagination.limit,
    });
  };

  const { data, mutate, error } = useSWR(
    `/feeds/${feedId}/articles?${queryString(pagination)}`,
    fetcher,
  );

  const loading = !data && !error;

  return {
    loading,
    error,
    articles: data?.articles,
    mutate,
  };
}
