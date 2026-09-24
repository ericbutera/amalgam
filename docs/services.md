# Services

## [Graph (GraphQL)](../services/graph/README.md)

[GraphQL API](http://localhost:8082). The goal is to show how to quickly build out public facing features. GraphQL would be available to public.

### [User Interface (UI)](../ui/README.md)

A Next.JS app [user interface](http://localhost:3000/) for interacting with the project. It uses the public GraphQL API.

## [RPC (gRPC)](../services/rpc/README.md)

A simple gRPC service that can be used to show how to convert a monolith into microservices architecture. This service would only be available on the internal VPC.

## Clients

- [Command Line Interface (CLI)](https://github.com/ericbutera/amalgam/tree/9528beb51c6b2affa3b6bd1622ca666983148fc4/cli)

## Data Pipeline

### Temporal

The first [pipeline](../data-pipeline/temporal/feed) that I have built is a rudimentary batch process for ingesting RSS feeds.

## Supporting Services

### LGTM Observability Stack

Observability is the heart of quality software. This project uses a demonstration LGTM stack to show how various pieces of the system can be monitored.

- [Grafana](http://localhost:3001/)
- [Prometheus](http://localhost:9090/)

TODO:

- [configure Loki](https://grafana.com/docs/loki/latest/)

### Minio

[Minio](https://min.io/) is a drop in object storage similar to AWS S3. I mainly will be using it as a data storage mechanism for data pipelines, but also as a fake CDN.

### [K6 (testing)](./k6/README.md)

K6 tests exercise the public GraphQL API. The smoke tests cover feed and
article reads, while the load and traffic tests live in the dedicated GraphQL
test directories.

## API boundary

- GraphQL is the public application API.
- gRPC is internal service communication and should not be publicly routed.
- The former REST API and its OpenAPI-generated clients are retired.
