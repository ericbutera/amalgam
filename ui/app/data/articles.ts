import useSWR from "swr";
// import { getApi } from '../lib/fetch';
import { getGraph } from '../lib/fetch';

export default function useArticles(id: string) {
  const fetcher = async () => await getGraph().ListArticles({feedId: id}) //await getApi().feedsIdArticlesGet({ id });
  const { data, mutate, error } = useSWR(`/feeds/${id}/articles`, fetcher);
  const loading = !data && !error;

  return {
    loading,
    error,
    articles: data?.articles,
    mutate
  };
}
