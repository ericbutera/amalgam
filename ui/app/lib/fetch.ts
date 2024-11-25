// TODO: use tanstack/react-query
import { GraphQLClient } from "graphql-request";
import { getSdk } from "../generated/graphql";

export const getGraph = (url?: string) => {
  // TODO: should this be a singleton?
  const endpoint = url ?? process.env.NEXT_PUBLIC_GRAPHQL_API_URL;

  if (!endpoint) {
    console.error("NEXT_PUBLIC_GRAPHQL_API_URL is not defined");
    throw new Error("NEXT_PUBLIC_GRAPHQL_API_URL is not defined");
  }

  return getSdk(new GraphQLClient(endpoint));
};
