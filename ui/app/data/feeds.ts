import useSWR from "swr";

import config from '../config.ts';
import fetch from '../lib/fetch';

export default function useFeeds() {
  const { data, mutate, error } = useSWR(`${config.apiHost}/feeds`, fetch);

  const loading = !data && !error;
  const feeds = data?.feeds || [];

  return {
    loading,
    error,
    feeds,
    mutate
  };
}
