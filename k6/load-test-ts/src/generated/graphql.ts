import { GraphQLClient, RequestOptions } from 'graphql-request';
import gql from 'graphql-tag';
export type Maybe<T> = T | null;
export type InputMaybe<T> = Maybe<T>;
export type Exact<T extends { [key: string]: unknown }> = { [K in keyof T]: T[K] };
export type MakeOptional<T, K extends keyof T> = Omit<T, K> & { [SubKey in K]?: Maybe<T[SubKey]> };
export type MakeMaybe<T, K extends keyof T> = Omit<T, K> & { [SubKey in K]: Maybe<T[SubKey]> };
export type MakeEmpty<T extends { [key: string]: unknown }, K extends keyof T> = { [_ in K]?: never };
export type Incremental<T> = T | { [P in keyof T]?: P extends ' $fragmentName' | '__typename' ? T[P] : never };
type GraphQLClientRequestHeaders = RequestOptions['requestHeaders'];
/** All built-in and custom scalars, mapped to their actual values */
export type Scalars = {
  ID: { input: string; output: string; }
  String: { input: string; output: string; }
  Boolean: { input: boolean; output: boolean; }
  Int: { input: number; output: number; }
  Float: { input: number; output: number; }
};

export type AddResponse = {
  __typename?: 'AddResponse';
  id: Scalars['ID']['output'];
};

export type Article = {
  __typename?: 'Article';
  authorEmail?: Maybe<Scalars['String']['output']>;
  authorName?: Maybe<Scalars['String']['output']>;
  content: Scalars['String']['output'];
  feedId: Scalars['ID']['output'];
  guid?: Maybe<Scalars['String']['output']>;
  id: Scalars['ID']['output'];
  imageUrl?: Maybe<Scalars['String']['output']>;
  preview: Scalars['String']['output'];
  title: Scalars['String']['output'];
  url: Scalars['String']['output'];
};

export type Feed = {
  __typename?: 'Feed';
  id: Scalars['ID']['output'];
  name: Scalars['String']['output'];
  url: Scalars['String']['output'];
};

export type Mutation = {
  __typename?: 'Mutation';
  addFeed: AddResponse;
  updateFeed: UpdateResponse;
};


export type MutationAddFeedArgs = {
  name: Scalars['String']['input'];
  url: Scalars['String']['input'];
};


export type MutationUpdateFeedArgs = {
  id: Scalars['ID']['input'];
  name?: InputMaybe<Scalars['String']['input']>;
  url?: InputMaybe<Scalars['String']['input']>;
};

export type Query = {
  __typename?: 'Query';
  article?: Maybe<Article>;
  articles: Array<Article>;
  feed?: Maybe<Feed>;
  feeds: Array<Feed>;
};


export type QueryArticleArgs = {
  id: Scalars['ID']['input'];
};


export type QueryArticlesArgs = {
  feedId: Scalars['ID']['input'];
};


export type QueryFeedArgs = {
  id: Scalars['ID']['input'];
};

export type UpdateResponse = {
  __typename?: 'UpdateResponse';
  id: Scalars['ID']['output'];
};

export type AddFeedMutationVariables = Exact<{
  url: Scalars['String']['input'];
  name: Scalars['String']['input'];
}>;


export type AddFeedMutation = { __typename?: 'Mutation', addFeed: { __typename?: 'AddResponse', id: string } };

export type UpdateFeedMutationVariables = Exact<{
  id: Scalars['ID']['input'];
  url?: InputMaybe<Scalars['String']['input']>;
  name?: InputMaybe<Scalars['String']['input']>;
}>;


export type UpdateFeedMutation = { __typename?: 'Mutation', updateFeed: { __typename?: 'UpdateResponse', id: string } };

export type ListFeedsQueryVariables = Exact<{ [key: string]: never; }>;


export type ListFeedsQuery = { __typename?: 'Query', feeds: Array<{ __typename?: 'Feed', id: string, url: string, name: string }> };

export type GetFeedQueryVariables = Exact<{
  id: Scalars['ID']['input'];
}>;


export type GetFeedQuery = { __typename?: 'Query', feed?: { __typename?: 'Feed', id: string, url: string, name: string } | null };

export type GetArticleQueryVariables = Exact<{
  id: Scalars['ID']['input'];
}>;


export type GetArticleQuery = { __typename?: 'Query', article?: { __typename?: 'Article', id: string, feedId: string, url: string, title: string, imageUrl?: string | null, content: string, preview: string, guid?: string | null, authorName?: string | null, authorEmail?: string | null } | null };

export type ListArticlesQueryVariables = Exact<{
  feedId: Scalars['ID']['input'];
}>;


export type ListArticlesQuery = { __typename?: 'Query', articles: Array<{ __typename?: 'Article', id: string, url: string, title: string, imageUrl?: string | null, preview: string, authorName?: string | null, authorEmail?: string | null }> };


export const AddFeedDocument = gql`
    mutation AddFeed($url: String!, $name: String!) {
  addFeed(url: $url, name: $name) {
    id
  }
}
    `;
export const UpdateFeedDocument = gql`
    mutation UpdateFeed($id: ID!, $url: String, $name: String) {
  updateFeed(id: $id, url: $url, name: $name) {
    id
  }
}
    `;
export const ListFeedsDocument = gql`
    query ListFeeds {
  feeds {
    id
    url
    name
  }
}
    `;
export const GetFeedDocument = gql`
    query GetFeed($id: ID!) {
  feed(id: $id) {
    id
    url
    name
  }
}
    `;
export const GetArticleDocument = gql`
    query GetArticle($id: ID!) {
  article(id: $id) {
    id
    feedId
    url
    title
    imageUrl
    content
    preview
    guid
    authorName
    authorEmail
  }
}
    `;
export const ListArticlesDocument = gql`
    query ListArticles($feedId: ID!) {
  articles(feedId: $feedId) {
    id
    url
    title
    imageUrl
    preview
    authorName
    authorEmail
  }
}
    `;

export type SdkFunctionWrapper = <T>(action: (requestHeaders?:Record<string, string>) => Promise<T>, operationName: string, operationType?: string, variables?: any) => Promise<T>;


const defaultWrapper: SdkFunctionWrapper = (action, _operationName, _operationType, _variables) => action();

export function getSdk(client: GraphQLClient, withWrapper: SdkFunctionWrapper = defaultWrapper) {
  return {
    AddFeed(variables: AddFeedMutationVariables, requestHeaders?: GraphQLClientRequestHeaders): Promise<AddFeedMutation> {
      return withWrapper((wrappedRequestHeaders) => client.request<AddFeedMutation>(AddFeedDocument, variables, {...requestHeaders, ...wrappedRequestHeaders}), 'AddFeed', 'mutation', variables);
    },
    UpdateFeed(variables: UpdateFeedMutationVariables, requestHeaders?: GraphQLClientRequestHeaders): Promise<UpdateFeedMutation> {
      return withWrapper((wrappedRequestHeaders) => client.request<UpdateFeedMutation>(UpdateFeedDocument, variables, {...requestHeaders, ...wrappedRequestHeaders}), 'UpdateFeed', 'mutation', variables);
    },
    ListFeeds(variables?: ListFeedsQueryVariables, requestHeaders?: GraphQLClientRequestHeaders): Promise<ListFeedsQuery> {
      return withWrapper((wrappedRequestHeaders) => client.request<ListFeedsQuery>(ListFeedsDocument, variables, {...requestHeaders, ...wrappedRequestHeaders}), 'ListFeeds', 'query', variables);
    },
    GetFeed(variables: GetFeedQueryVariables, requestHeaders?: GraphQLClientRequestHeaders): Promise<GetFeedQuery> {
      return withWrapper((wrappedRequestHeaders) => client.request<GetFeedQuery>(GetFeedDocument, variables, {...requestHeaders, ...wrappedRequestHeaders}), 'GetFeed', 'query', variables);
    },
    GetArticle(variables: GetArticleQueryVariables, requestHeaders?: GraphQLClientRequestHeaders): Promise<GetArticleQuery> {
      return withWrapper((wrappedRequestHeaders) => client.request<GetArticleQuery>(GetArticleDocument, variables, {...requestHeaders, ...wrappedRequestHeaders}), 'GetArticle', 'query', variables);
    },
    ListArticles(variables: ListArticlesQueryVariables, requestHeaders?: GraphQLClientRequestHeaders): Promise<ListArticlesQuery> {
      return withWrapper((wrappedRequestHeaders) => client.request<ListArticlesQuery>(ListArticlesDocument, variables, {...requestHeaders, ...wrappedRequestHeaders}), 'ListArticles', 'query', variables);
    }
  };
}
export type Sdk = ReturnType<typeof getSdk>;
