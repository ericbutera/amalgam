# Services

## [Graph (GraphQL)](./graph/README.md)

[GraphQL API](http://localhost:8082). The goal is to show how to quickly build out public facing features. GraphQL would be available to public.

### [User Interface (UI)](./ui/README.md)

A Next.JS app [user interface](http://localhost:3000/) for interacting with the project. At first it uses the REST API, but will be updated to use the GraphQL API.

## [RPC (gRPC)](./rpc/README.md)

A simple gRPC service that can be used to show how to convert a monolith into microservices architecture. This service would only be available on the internal VPC.

## Clients

- [Command Line Interface (CLI)](https://github.com/ericbutera/amalgam/tree/9528beb51c6b2affa3b6bd1622ca666983148fc4/cli)

## Data Pipeline

### Temporal

The first [pipeline](./data-pipeline/temporal/feed) that I have built is a rudimentary batch process for ingesting RSS feeds.

## Supporting Services

### [LGTM Observability Stack](./lgtm/README.md)

Observability is the heart of quality software. This project uses a demonstration LGTM stack to show how various pieces of the system can be monitored.

- [Grafana](http://localhost:3001/)
- [Prometheus](http://localhost:9090/)
- [API Dashboard](http://localhost:3001/d/amalgam-gin-dashboard/gin-application-metrics?orgId=1&refresh=5s)
- [Amalgam Dashboard](http://localhost:3001/d/amalgam-dashboard/amalgam?orgId=1&refresh=5s)

TODO:

- [configure Loki](https://grafana.com/docs/loki/latest/)

### Minio

[Minio](https://min.io/) is a drop in object storage similar to AWS S3. I mainly will be using it as a data storage mechanism for data pipelines, but also as a fake CDN.

### [K6 (testing)](./k6/README.md)

K6 tests have been [generated](./k6/tests/README.md) from the OpenAPI spec. They are a high level way of verifying the API is working as expected. This is a wonderful way to have end-to-end tests that are easy to write and maintain. Next steps might be adding load testing.

## Deprecated Services

### ~~[Monolith REST API](./api/README.md)~~ deprecated [v1.4.0](https://github.com/ericbutera/amalgam/releases/tag/v1.4.0)

A classic "monolith" REST api built with gin-gonic. This service would be available to the public. It will be replaced with GraphQL as the project progresses.

One of the major points of this project is the OpenAPI specification. It is generated from the gin endpoint code and is used to generate API clients in various languages.

- [API](http://localhost:8080)
- [metrics](http://localhost:8080/metrics)
