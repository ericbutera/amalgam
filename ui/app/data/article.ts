import useSWR from "swr";
// import { getApi } from '../lib/fetch';
import { getGraph } from '../lib/fetch';

export default function useArticle(id: string) {
  const fetcher = async () => await getGraph().GetArticle({ id }) //getApi().articlesIdGet({ id });
  const { data, mutate, error } = useSWR(`/article/${id}`, fetcher);
  const loading = !data && !error;

  return {
    loading,
    error,
    article: data?.article,
    mutate
  };
}
