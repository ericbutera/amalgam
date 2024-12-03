import useSWR from "swr";
import { getGraph } from "../lib/fetch";

export default function useArticles(feedId: string, cursor?: string) {
  const fetcher = async () => {
    return await getGraph().ListArticles({ feedId, cursor });
  };

  const { data, mutate, error } = useSWR(
    cursor
      ? `/feeds/${feedId}/articles?cursor=${cursor}`
      : `/feeds/${feedId}/articles`,
    fetcher
  );

  const loading = !data && !error;

  return {
    loading,
    error,
    articles: data?.articles,
    mutate,
  };
}
