// TODO: use tanstack/react-query
import { GraphQLClient } from 'graphql-request';
import { getSdk } from '../generated/graphql';

export const getGraph = (url: string) => {
  if (!process.env.NEXT_PUBLIC_GRAPHQL_API_URL) {
    throw new Error('NEXT_PUBLIC_GRAPHQL_API_URL is not defined');
  }
  // TODO: should this be a singleton?
  const endpoint = url || process.env.NEXT_PUBLIC_GRAPHQL_API_URL;
  return getSdk(new GraphQLClient(endpoint));
}
