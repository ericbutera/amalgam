import useSWR from "swr";
// import { getApi } from '../lib/fetch'; // TODO: delete getApi
import { getGraph } from '../lib/fetch';

export default function useFeeds() {
  const fetcher = async () => {
    return await getGraph().ListFeeds();
    //return await getApi().feedsGet();
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
