import { DefaultApi } from '../lib/client/apis/DefaultApi';

export default async function fetcher<JSON = any>(
  input: RequestInfo,
  init?: RequestInit
): Promise<JSON> {
  const res = await fetch(input, init)
  return res.json()
}

// TODO: DefaultApi is probably the wrong usage of this; needs more research
export const getApi = () => new DefaultApi();