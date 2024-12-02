import useSWR from "swr";
import { getGraph } from "../lib/fetch";
import toast from "react-hot-toast";

export default function useArticle(id: string) {
  const fetcher = async () => {
    const graph = getGraph();
    const article = await graph.GetArticle({ id });

    if (article && !article.article?.userArticle?.viewedAt) {
      graph.MarkArticleRead({ id }).catch(() => {
        toast.error("Failed to mark article as read");
      });
    }

    return article;
  };
  const { data, mutate, error } = useSWR(`/article/${id}`, fetcher);
  const loading = !data && !error;

  return {
    loading,
    error,
    article: data?.article,
    mutate,
  };
}
