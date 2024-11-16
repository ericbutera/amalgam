# Amalgam

Tech demo of a modern web application stack.

## Goals

- Developer Experience: Spin up the entire stack with a single command.
- Local-First Development: Everything runs locally for true production parity—no surprises.
- Integrated Observability: Built-in o11y from the start, not an afterthought.

## Technologies

- Kubernetes: The backbone for modern orchestration.
- GraphQL: A powerful API layer for flexible, efficient data fetching.
- Golang: Performance and simplicity at scale.
- Temporal: Resilient workflows made easy.
- Grafana LGTM Stack: Observe, debug, and improve with confidence.
- Next.js + TypeScript: A delightful developer experience for building blazing-fast, modern UIs.
- MySQL: The trusted relational database powering countless applications.

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

[Code Quality](./docs/code-quality.md) contains an overview of how to ensure the codebase is of high quality.

## CI/CD

Github Actions can be found in the [.github/workflows](./.github/workflows) directory. You can run them locally using [act](https://github.com/nektos/act).
