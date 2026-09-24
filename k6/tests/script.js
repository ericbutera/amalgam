import http from 'k6/http';
import { check } from 'k6';

export const options = {
  vus: 1,
  iterations: 1,
};

const GRAPH_HOST = __ENV.GRAPH_HOST || 'http://localhost:8082/query';
const headers = { headers: { 'Content-Type': 'application/json' } };

export default function () {
  const feedsResponse = post({ query: `
    query ListFeeds {
      feeds { feeds { id url name } }
    }
  ` });
  check(feedsResponse, {
    'list feeds returns HTTP 200': (response) => response.status === 200,
    'list feeds returns GraphQL data': (response) => response.json('data.feeds.feeds') !== null,
  });

  const feedID = feedsResponse.json('data.feeds.feeds.0.id');
  if (!feedID) {
    return;
  }

  const feedResponse = post({
    query: `query GetFeed($id: ID!) { feed(id: $id) { id url name } }`,
    variables: { id: feedID },
  });
  check(feedResponse, {
    'get feed returns HTTP 200': (response) => response.status === 200,
    'get feed returns the selected feed': (response) => response.json('data.feed.id') === feedID,
  });

  const articlesResponse = post({
    query: `query ListArticles($feedId: ID!) {
      articles(feedId: $feedId) { articles { id } }
    }`,
    variables: { feedId: feedID },
  });
  check(articlesResponse, {
    'list articles returns HTTP 200': (response) => response.status === 200,
    'list articles returns GraphQL data': (response) => response.json('data.articles') !== null,
  });

  const articleID = articlesResponse.json('data.articles.articles.0.id');
  if (!articleID) {
    return;
  }

  const articleResponse = post({
    query: `query GetArticle($id: ID!) { article(id: $id) { id feedId } }`,
    variables: { id: articleID },
  });
  check(articleResponse, {
    'get article returns HTTP 200': (response) => response.status === 200,
    'get article returns the selected article': (response) => response.json('data.article.id') === articleID,
  });
}

function post(body) {
  return http.post(GRAPH_HOST, JSON.stringify(body), headers);
}
