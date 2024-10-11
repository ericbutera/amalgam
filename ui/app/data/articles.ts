import useSWR from "swr";

import config from '../config.ts';
import fetch from '../lib/fetch';

export default function useArticles(id) {
  const { data, mutate, error } = useSWR(`${config.apiHost}/feed/${id}/articles`, fetch);

  const loading = !data && !error;
  const articles = data?.articles || [];

  return {
    loading,
    error,
    articles,
    mutate
  };
}
