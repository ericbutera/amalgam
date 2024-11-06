import useSWR from "swr";
import { getGraph } from '../lib/fetch';

export default function useFeeds() {
  const fetcher = async () => {
    return await getGraph().ListFeeds();
  }
  const { data, mutate, error } = useSWR(`/feeds`, fetcher);
  const loading = !data && !error;

  if (error) {
    // todo: splash
    console.error(error);
  }

  return {
    loading,
    error,
    feeds: data?.feeds,
    mutate
  };
}
