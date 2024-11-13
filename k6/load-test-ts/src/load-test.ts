import './fetch-polyfill'

import { group, check, sleep } from 'k6';

import { GraphQLClient } from 'graphql-request';
import { getSdk, ListFeedsQuery } from './generated/graphql';
const getGraph = (url: string='') => {
  const endpoint = url || __ENV.GRAPH_HOST;
  if (!endpoint) {
    throw new Error('GRAPH_HOST is not defined');
  }
  return getSdk(new GraphQLClient(endpoint));
}

if (!globalThis.fetch) {
  globalThis.fetch = fetch
  //globalThis.Headers = new fetch.Headers();
}


const graph = getGraph();

export default async function () {
  const feeds = await graph.ListFeeds();
  check(feeds, { 'feeds is not null': (r) => r.feeds.length > 0 })
}

/*
import { getSdk, ListFeedsDocument, ListFeedsQuery } from './generated/graphql';

import {  check, } from 'k6';
import http from 'k6/http';
import { DocumentNode } from 'graphql';

const endpoint = __ENV.GRAPH_HOST || 'http://localhost:8082/query';

const headers = {
  'Content-Type': 'application/json',
};

export default async function () {
  const data = getParams(ListFeedsDocument {})
  console.log("data", data)
  const res = http.post(
    endpoint,
    data,
    { headers }
  );
  check(res, { 'status code is 200': (r) => r.status === 200 });
  // check(res.feeds, { 'feeds is not empty': (r) => r.feeds.length > 0 });
}

const getParams = (query: DocumentNode, variables: any) => {
  return JSON.stringify({
    query.toString(),
    variables
  });
}
*/

/*
// export let options = {
//   stages: [
//     { duration: '30s', target: 10 }, // Ramp up to 10 users
//     { duration: '1m', target: 10 },  // Sustain 10 users
//     { duration: '30s', target: 0 },  // Ramp down to 0 users
//   ],
// };

// const feeds = await graph.ListFeeds();
// const articles = graph.ListArticles({ feedId: id })

  if (feeds.feeds.length > 0) {
    // TODO:
    //let feedId = feeds.feeds[0].id;

    // GET /feeds/:id
    // let feedDetailsResponse = http.get(`${BASE_URL}/feeds/${feedId}`);
    // check(feedDetailsResponse, { 'status is 200': (r) => r.status === 200 });
    // TODO:
    // const feed = await graph.GetFeed({ id: feedId })

    // GET /feeds/:id/articles
    // let feedArticlesResponse = http.get(`${BASE_URL}/feeds/${feedId}/articles`);
    // check(feedArticlesResponse, { 'status is 200': (r) => r.status === 200 });
    // TODO
    // const articles = graph.ListArticles({ feedId: feedId.id })

    // Assuming `feedArticlesResponse` returns a list of articles with IDs
    // let articles = feedArticlesResponse.json();
    // if (articles.length > 0) {
    //   let articleId = articles[0].id;
    //   // GET /articles/:id
    //   let articleResponse = http.get(`${BASE_URL}/articles/${articleId}`);
    //   check(articleResponse, { 'status is 200': (r) => r.status === 200 });
    // }

    if (articles.articles.length > 0) {
      let articleId = articles.articles[0].id;
      // GET /articles/:id
      graph.ListArticles({ feedId: feedId })
    }
      */
