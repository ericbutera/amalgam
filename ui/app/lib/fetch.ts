export default async function fetcher<JSON = any>(
  input: RequestInfo,
  init?: RequestInit
): Promise<JSON> {
  const res = await fetch(input, init)
  return res.json()
}

// replace with getGraph
import { GraphQLClient } from 'graphql-request';
import { getSdk } from '../generated/graphql';

// TODO: tanstack/react-query
// TODO: url configuration, singleton
const client = new GraphQLClient('http://localhost:8082/query');
const sdk = getSdk(client);

export const getGraph = () => sdk;
