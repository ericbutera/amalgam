import useSWR from "swr";
import { getApi } from '../lib/fetch';

export default function useFeeds() {
  const fetcher = async () => await getApi().feedsGet();
  const { data, mutate, error } = useSWR(`/feeds`, fetcher);
  const loading = !data && !error;

  return {
    loading,
    error,
    feeds: data?.feeds,
    mutate
  };
}
