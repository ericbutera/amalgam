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

A REST api built with gin-gonic.

- [API](http://localhost:8080)
- [metrics](http://localhost:8080/metrics)

### User Interface (UI)

A Next.JS app.

- [UI](http://localhost:3000/)

### LGTM

Observability stack.

- [Grafana](http://localhost:3001/)
- [Prometheus](http://localhost:9090/)

## Code generation

- [REST client](./pkg/client/README.md) from OpenAPI spec
- [k6 tests](./k6/README.md) from OpenAPI spec
