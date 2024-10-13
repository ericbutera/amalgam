import useSWR from "swr";
import { getApi } from '../lib/fetch';

export default function useArticles(id: number) {
  const fetcher = async () => await getApi().feedsIdArticlesGet({ id });
  const { data, mutate, error } = useSWR(`/feeds/${id}/articles`, fetcher);
  const loading = !data && !error;

  return {
    loading,
    error,
    articles: data?.articles,
    mutate
  };
}
