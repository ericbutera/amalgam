/* eslint-disable */
import * as types from './graphql';
import { TypedDocumentNode as DocumentNode } from '@graphql-typed-document-node/core';

/**
 * Map of all GraphQL operations in the project.
 *
 * This map has several performance disadvantages:
 * 1. It is not tree-shakeable, so it will include all operations in the project.
 * 2. It is not minifiable, so the string of a GraphQL query will be multiple times inside the bundle.
 * 3. It does not support dead code elimination, so it will add unused operations.
 *
 * Therefore it is highly recommended to use the babel or swc plugin for production.
 * Learn more about it here: https://the-guild.dev/graphql/codegen/plugins/presets/preset-client#reducing-bundle-size
 */
const documents = {
    "mutation AddFeed($url: String!, $name: String!) {\n  addFeed(url: $url, name: $name) {\n    id\n  }\n}\n\nmutation UpdateFeed($id: ID!, $url: String, $name: String) {\n  updateFeed(id: $id, url: $url, name: $name) {\n    id\n  }\n}": types.AddFeedDocument,
    "query ListFeeds {\n  feeds {\n    id\n    url\n    name\n  }\n}\n\nquery GetFeed($id: ID!) {\n  feed(id: $id) {\n    id\n    url\n    name\n  }\n}\n\nquery GetArticle($id: ID!) {\n  article(id: $id) {\n    id\n    feedId\n    url\n    title\n    imageUrl\n    content\n    preview\n    guid\n    authorName\n    authorEmail\n  }\n}\n\nquery ListArticles($feedId: ID!) {\n  articles(feedId: $feedId) {\n    id\n    url\n    title\n    imageUrl\n    preview\n    authorName\n    authorEmail\n  }\n}": types.ListFeedsDocument,
};

/**
 * The graphql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 *
 *
 * @example
 * ```ts
 * const query = graphql(`query GetUser($id: ID!) { user(id: $id) { name } }`);
 * ```
 *
 * The query argument is unknown!
 * Please regenerate the types.
 */
export function graphql(source: string): unknown;

/**
 * The graphql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function graphql(source: "mutation AddFeed($url: String!, $name: String!) {\n  addFeed(url: $url, name: $name) {\n    id\n  }\n}\n\nmutation UpdateFeed($id: ID!, $url: String, $name: String) {\n  updateFeed(id: $id, url: $url, name: $name) {\n    id\n  }\n}"): (typeof documents)["mutation AddFeed($url: String!, $name: String!) {\n  addFeed(url: $url, name: $name) {\n    id\n  }\n}\n\nmutation UpdateFeed($id: ID!, $url: String, $name: String) {\n  updateFeed(id: $id, url: $url, name: $name) {\n    id\n  }\n}"];
/**
 * The graphql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function graphql(source: "query ListFeeds {\n  feeds {\n    id\n    url\n    name\n  }\n}\n\nquery GetFeed($id: ID!) {\n  feed(id: $id) {\n    id\n    url\n    name\n  }\n}\n\nquery GetArticle($id: ID!) {\n  article(id: $id) {\n    id\n    feedId\n    url\n    title\n    imageUrl\n    content\n    preview\n    guid\n    authorName\n    authorEmail\n  }\n}\n\nquery ListArticles($feedId: ID!) {\n  articles(feedId: $feedId) {\n    id\n    url\n    title\n    imageUrl\n    preview\n    authorName\n    authorEmail\n  }\n}"): (typeof documents)["query ListFeeds {\n  feeds {\n    id\n    url\n    name\n  }\n}\n\nquery GetFeed($id: ID!) {\n  feed(id: $id) {\n    id\n    url\n    name\n  }\n}\n\nquery GetArticle($id: ID!) {\n  article(id: $id) {\n    id\n    feedId\n    url\n    title\n    imageUrl\n    content\n    preview\n    guid\n    authorName\n    authorEmail\n  }\n}\n\nquery ListArticles($feedId: ID!) {\n  articles(feedId: $feedId) {\n    id\n    url\n    title\n    imageUrl\n    preview\n    authorName\n    authorEmail\n  }\n}"];

export function graphql(source: string) {
  return (documents as any)[source] ?? {};
}

export type DocumentType<TDocumentNode extends DocumentNode<any, any>> = TDocumentNode extends DocumentNode<  infer TType,  any>  ? TType  : never;
