# Amalgam

Tech demo of a modern web application stack.

## Prerequisites

1. docker
2. tilt.dev

## Run

```sh
tilt up
curl -v http://localhost:3000/ # or open in browser
```

## Services

### [Monolith REST API](./api/README.md)

A classic "monolith" REST api built with gin-gonic. This service would be available to the public. It will be replaced with GraphQL as the project progresses.

One of the major points of this project is the OpenAPI specification. It is generated from the gin endpoint code and is used to generate API clients in various languages.

- [API](http://localhost:8080)
- [metrics](http://localhost:8080/metrics)

### [Graph (GraphQL)](./graph/README.md)

[GraphQL API](http://localhost:8082). The goal is to show how to quickly build out public facing features. GraphQL would be available to public.

Resources:

- [tools and libraries](https://graphql.org/community/tools-and-libraries/?tags=go)

### [User Interface (UI)](./ui/README.md)

A Next.JS app [user interface](http://localhost:3000/) for interacting with the project. At first it uses the REST API, but will be updated to use the GraphQL API.

TODO:

- [apollo-client](https://github.com/apollographql/apollo-client)
- [relay](https://github.com/facebook/relay)

### [Command Line Interface (CLI)](./cli/README.md)

A way to interact with the project using the command line. This project aims to be low-code as it uses a pre-generated REST API client. A stretch goal will be to add a Text User Interface (TUI).

### [RPC (gRPC)](./rpc/README.md)

A simple gRPC service that can be used to show how to convert a monolith into microservices architecture. This service would only be available on the internal VPC.

### Data Pipeline

TODO: ingest various RSS feeds using different technologies and strategies.

### [LGTM Observability Stack](./lgtm/README.md)

Observability is the heart of quality software. This project uses a demonstration LGTM stack to show how various pieces of the system can be monitored.

- [Grafana](http://localhost:3001/)
- [Prometheus](http://localhost:9090/)
- [API Dashboard](http://localhost:3001/d/amalgam-gin-dashboard/gin-application-metrics?orgId=1&refresh=5s)
- [Amalgam Dashboard](http://localhost:3001/d/amalgam-dashboard/amalgam?orgId=1&refresh=5s)

TODO:

- [configure Loki](https://grafana.com/docs/loki/latest/)

### [K6 (testing)](./k6/README.md)

K6 tests have been [generated](./k6/tests/README.md) from the OpenAPI spec. They are a high level way of verifying the API is working as expected. This is a wonderful way to have end-to-end tests that are easy to write and maintain. Next steps might be adding load testing.

## Domain Logic

A simplified version of domain driven design.

- [service](./internal/service/service.go)

TODO:

- [validation](https://github.com/go-playground/validator/blob/master/_examples/struct-level/main.go)

## Code generation

One of the major goals of this project is to show how to quickly build integrations. Part of that is utilizing code generation to lower the amount of code that needs to be written and maintained.

- [OpenAPI spec](./api/docs/swagger.yaml) with [swaggo/swag](https://github.com/swaggo/swag)
- [REST client](./pkg/client/README.md) from OpenAPI spec
- [TypeScript client](./ui/app/lib/client/) from OpenAPI spec
- [k6 tests](./k6/README.md) from OpenAPI spec

## TODO

- helm chart
  - values file for configuration
  - service account
  - ingress
  - deployment + service
  - base chart to share otel conf
- [just](https://github.com/casey/just)
- ~~graphql~~
- fake feed data generation
- [ci gate: coverage](https://github.com/vladopajic/go-test-coverage)
