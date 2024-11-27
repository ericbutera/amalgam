import useSWR from "swr";
import { getGraph } from "../lib/fetch";

export default function useFeed(id: string) {
  const fetcher = async () => await getGraph().GetFeed({ id });
  const { data, mutate, error } = useSWR(id ? `/feed/${id}` : null, fetcher);
  const loading = !data && !error;

  return {
    loading,
    error,
    feed: data?.feed,
    mutate,
  };
}
