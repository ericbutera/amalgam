# Tests

This demo shows how to generate k6 tests from an OpenAPI spec. These tests can be used to validate & load test the API.

Read the [Grafana K6 docs](https://grafana.com/docs/k6/latest/) for more information.

## Quick start

```sh
# prereq: generate ./api/docs/swagger.json
make generate
make test
```

## TODO

- [fake data generation](https://github.com/grafana/k6-example-data-generation/blob/main/src/index.js)
- [load test](https://k6.io/blog/load-testing-your-api-with-swagger-openapi-and-k6/)
