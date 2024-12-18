import useSWR from "swr";
import { getGraph } from "../lib/fetch";

export default function useFeeds() {
  const fetcher = async () => await getGraph().ListFeeds();
  const { data, mutate, error } = useSWR(`/feeds`, fetcher);
  const loading = !data && !error;

  return {
    loading,
    error,
    feeds: data?.feeds?.feeds,
    mutate,
  };
}
