# Amalgam

Tech demo.

## Prerequisites

1. docker
2. tilt.dev

## Run

```sh
tilt up
curl -v http://localhost:3000/ # or open in browser
```

## Services

### API

A classic REST api built with gin-gonic.

- [API](http://localhost:8080)
- [metrics](http://localhost:8080/metrics)

### Graph

[GraphQL API](http://localhost:8082). The goal is to show how to quickly build out public facing features with GraphqL using data gathered from data pipelines.

### User Interface (UI)

A Next.JS app [user interface](http://localhost:3000/).

### LGTM

Observability stack.

- [Grafana](http://localhost:3001/)
- [Prometheus](http://localhost:9090/)
- [API Dashboard](http://localhost:3001/d/amalgam-gin-dashboard/gin-application-metrics?orgId=1&refresh=5s)
- [Amalgam Dashboard](http://localhost:3001/d/amalgam-dashboard/amalgam?orgId=1&refresh=5s)

### K6 (testing)

K6 tests have been [generated](./k6/tests/README.md) from the OpenAPI spec. They are a high level way of verifying the API is working as expected. This is a wonderful way to have end-to-end tests that are easy to write and maintain. Next steps might be adding load testing.

### Data Pipeline

TODO: ingest various RSS feeds using different technologies and strategies.

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
- graphql
- data pipeline demos
- fake feed data generation
- move api/docs to static file server
  - research: <https://github.com/scalar/scalar>
