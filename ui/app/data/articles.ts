import useSWR from "swr";
import { getGraph } from "../lib/fetch";

export default function useArticles(feedId: string) {
  const fetcher = async () => await getGraph().ListArticles({ feedId });
  const { data, mutate, error } = useSWR(`/feeds/${feedId}/articles`, fetcher);
  const loading = !data && !error;

  return {
    loading,
    error,
    articles: data?.articles,
    mutate,
  };
}
