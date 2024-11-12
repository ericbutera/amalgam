// https://k6.io/blog/compare-rest-and-graphql-using-k6-for-performance-testing/

import http from 'k6/http';
import { check, sleep } from 'k6';
// import { uuidv4 } from 'https://jslib.k6.io/k6-utils/1.4.0/index.js';

export let options = {
  stages: [
    { duration: '10s', target: 1_000 },
    { duration: '20s', target: 4_000 },
    { duration: '30s', target: 0 },
  ],
};

const BASE_URL = __ENV.GRAPH_HOST;

if (!BASE_URL) {
  throw new Error('GRAPH_HOST is not defined')
}

export default function () {

  let queryRes = http.post(BASE_URL, JSON.stringify({ query: QueryListFeeds }), {
    headers: { 'Content-Type': 'application/json' },
  });
  check(queryRes, { 'query status is 200': (r) => r.status === 200 });

  // const variables = { url: `http://faker:8084/feed/${uuidv4()}` };
  // let res = http.post(BASE_URL, JSON.stringify({ query: MutationAddFeed, variables: variables }), {
  //   headers: { 'Content-Type': 'application/json' },
  // });
  // check(res, { 'status 200': (r) => r.status === 200 })

  sleep(1);
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
