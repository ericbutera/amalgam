import useSWR from "swr";

import config from '../config.ts';
import fetch from '../lib/fetch';

export default function useArticle(id) {
  const { data, mutate, error } = useSWR(`${config.apiHost}/article/${id}`, fetch);

  const loading = !data && !error;
  const article = data?.article || {};

  return {
    loading,
    error,
    article,
    mutate
  };
}

