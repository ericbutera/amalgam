import useSWR from "swr";
import { getGraph } from "../lib/fetch";

export default function useArticle(id: string) {
  const fetcher = async () => await getGraph().GetArticle({ id });
  const { data, mutate, error } = useSWR(`/article/${id}`, fetcher);
  const loading = !data && !error;

  return {
    loading,
    error,
    article: data?.article,
    mutate,
  };
}
