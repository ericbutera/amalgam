// https://k6.io/blog/compare-rest-and-graphql-using-k6-for-performance-testing/

import http from 'k6/http';
import { check, sleep } from 'k6';
import { uuidv4 } from 'https://jslib.k6.io/k6-utils/1.4.0/index.js';
import { randomString } from 'https://jslib.k6.io/k6-utils/1.4.0/index.js';

export let options = {
  stages: [
    { duration: '10s', target: 1 },
    { duration: '10s', target: 10 },
    { duration: '10s', target: 0 },
  ],
};

const GRAPH_HOST = __ENV.GRAPH_HOST;
const FAKER_HOST = __ENV.FAKER_HOST;

if (!GRAPH_HOST) {
  throw new Error('GRAPH_HOST is not defined')
}
if (!FAKER_HOST) {
  throw new Error('FAKER_HOST is not defined')
}

const DefaultHeaders = {
  headers: { 'Content-Type': 'application/json' },
};

export default function () {
  const feedsRes = post(listFeeds())
  check(feedsRes, { 'List Feeds status is 200': (r) => r.status === 200 });

  let feeds = feedsRes.json('data.feeds.#.id') || [];
  for (let i = 0; i < feeds.length; i++) {
    const feedId = feeds[i];
    if (!feedId)
      return

    check(post(getFeed(feedId)), { 'Get Feed status is 200': (r) => r.status === 200 });

    const articlesRes = post(listArticles(feedId));
    check(articlesRes, { 'List Articles status is 200': (r) => r.status === 200 });

    const articleId = articlesRes ? articlesRes.json('data.articles[0].id') : null; // TODO: randomize
    if (!articleId)
      return

    check(post(getArticle(articleId)), { 'Get Article status is 200': (r) => r.status === 200 });
  }
}

function post(body, headers) {
  headers = headers || DefaultHeaders;
  return http.post(GRAPH_HOST, JSON.stringify(body), headers);
}

function listFeeds() {
  return { query: QueryListFeeds };
}

function getFeed(feedId) {
  return { query: QueryGetFeed, variables: { id: feedId } };
}

function addFeed() {
  const feedId = uuidv4();
  const variables = {
    url: `http://${FAKER_HOST}/feed/${feedId}`,
    name: randomString(10, 'abcdefghijklmnopqrstuvwxyz '),
  };
  return { query: MutationAddFeed, variables };
}

function listArticles(feedId) {
  return { query: QueryListArticles, variables: { feedId: feedId } };
}

function getArticle(articleId) {
  return { query: QueryGetArticle, variables: { id: articleId } };
}

const QueryListFeeds = `
  query Feeds {
    feeds {
      id
      url
      name
    }
  }
`;
const MutationAddFeed = `
  mutation AddFeed($url: String!, $name: String!) {
    addFeed(url: $url, name: $name) {
      id
    }
  }
`;

const QueryGetFeed = `
  query GetFeed($id: ID!) {
    feed(id: $id) {
      id
      url
      name
    }
  }
`

const QueryGetArticle = `
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
`

const QueryListArticles = `
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
`
