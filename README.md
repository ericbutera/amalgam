# Amalgam

Tech demo of a modern web application stack.

## Goals

- Developer Experience
- Run everything locally (no surprises in production)
- o11y during development (not tacked on later)

## Technologies

- Kubernetes
- GraphQL
- Golang
- Temporal
- Grafana LGTM Observability
- Next.js + TypeScript
- MySQL

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

## Code Generation

One of the major goals of this project is to show how to quickly build integrations. Part of that is utilizing code generation to lower the amount of code that needs to be written and maintained. More information can be found in the [Code Generation](./docs/code-generation.md) document.

## Testing

[Testing](./docs/testing.md) contains an overview of how to test the various components of the application.

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
