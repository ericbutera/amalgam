import useSWR from "swr";
import { getGraph } from '../lib/fetch';

export default function useArticles(id: string) {
  const fetcher = async () => await getGraph().ListArticles({feedId: id})
  const { data, mutate, error } = useSWR(`/feeds/${id}/articles`, fetcher);
  const loading = !data && !error;

  return {
    loading,
    error,
    articles: data?.articles,
    mutate
  };
}
