import useSWR from "swr";
import { getApi } from '../lib/fetch';

export default function useArticle(id: number) {
  const fetcher = async () => await getApi().articlesIdGet({ id });
  const { data, mutate, error } = useSWR(`/article/${id}`, fetcher);
  const loading = !data && !error;

  return {
    loading,
    error,
    article: data?.article,
    mutate
  };
}

