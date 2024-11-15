# Amalgam

Tech demo of a modern web application stack.

## Prerequisites

1. [docker](https://docs.docker.com/get-docker/)
2. [tilt.dev](https://tilt.dev/)

## Run

```sh
tilt up
# open tilt ui @ http://localhost:10350
```

## Architecture

[Architecture](./docs/architecture.md) contains an overview of how to convert a project from a monolith to a microservices architecture.

## Code generation

One of the major goals of this project is to show how to quickly build integrations. Part of that is utilizing code generation to lower the amount of code that needs to be written and maintained.

- GraphQL
  - [TypeScript](https://github.com/ericbutera/amalgam/blob/9528beb51c6b2affa3b6bd1622ca666983148fc4/ui/app/generated/graphql.ts#L204-L225)
  - [Go](https://github.com/ericbutera/amalgam/blob/9528beb51c6b2affa3b6bd1622ca666983148fc4/pkg/clients/graphql/graphql.gen.go)
- OpenAPI
  - [OpenAPI spec](./api/docs/swagger.yaml) with [swaggo/swag](https://github.com/swaggo/swag)
  - [REST client](./pkg/client/README.md) from OpenAPI spec
  - [TypeScript client](./ui/app/lib/client/) from OpenAPI spec
  - [k6 tests](./k6/README.md) from OpenAPI spec

## Code Quality

Be sure to install the pre-commit hooks which run various linters, formatters, and tests.

```sh
just setup
```

Linters:

- [golangci-lint](https://golangci-lint.run/)
- [eslint](https://eslint.org/)
- [conventional commits](https://www.conventionalcommits.org/en/v1.0.0/)

## CI/CD

Github Actions can be found in the [.github/workflows](./.github/workflows) directory. You can run them locally using [act](https://github.com/nektos/act).
