// Simulate a small amount of traffic, much like a user would interact with the application.
// This is useful for testing out the observability stack.

import http from 'k6/http';
import { check, sleep } from 'k6';
//import { sample } from 'k6/vendor/ramda';
// import * as lodash from 'lodash';

export let options = {
  scenarios: {
    constant_load: {
      executor: 'constant-vus', // keeps a constant number of virtual users
      vus: 1,                   // adjust the number of VUs (virtual users) as needed
      duration: '87600h',       // run for 10 years (essentially infinite)
    },
  },
};

const GRAPH_HOST = __ENV.GRAPH_HOST;

if (!GRAPH_HOST) {
  throw new Error('GRAPH_HOST is not defined')
}

const DefaultHeaders = {
  headers: { 'Content-Type': 'application/json' },
};

export default function () {
  const feedsRes = post(listFeeds())
  check(feedsRes, { 'List Feeds status is 200': (r) => r.status === 200 });

  let feeds = feedsRes.json('data.feeds.#.id') || [];
  for (let feedId of feeds) {
    check(post(getFeed(feedId)), { 'Get Feed status is 200': (r) => r.status === 200 });

    const articlesRes = post(listArticles(feedId));
    check(articlesRes, { 'List Articles status is 200': (r) => r.status === 200 });

    const articles = articlesRes.json('data.articles.#.id') || [];
    let counter = 0
    for (let articleId of articles) {
      check(post(getArticle(articleId)), { 'Get Article status is 200': (r) => r.status === 200 });
      counter++;
      if (counter == 5) {
        sleep(1)
        counter = 0
      }
    }

    sleep(1)
  }

  sleep(1)
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
      description
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
