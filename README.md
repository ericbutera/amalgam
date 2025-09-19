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

Install the required prerequisites:

1. [docker](https://docs.docker.com/get-docker/) or [orbstack](https://www.orbstack.com/)
2. [asdf-vm](https://asdf-vm.com/)

Run `asdf install` to install additional required tools.

One caveat is that you may need to add individual asdf [plugins](https://asdf-vm.com/manage/plugins.html). Due to security reasons, asdf doesn't provide a native way to automatically add missing plugins during install. If you are comfortable with the tools listed in [.tool-versions](./.tool-versions), you can run the following command: `cut -d' ' -f1 .tool-versions | xargs -I {} asdf plugin add {}`, then `asdf install`. You can verify the plugins are installed by running `asdf plugin list`.

## Run

```sh
tilt up
# open tilt ui @ http://localhost:10350
```

## Development

[Development](./docs/development.md) contains an overview of how to develop the various components of the application.

## Architecture

[Architecture](./docs/architecture.md) contains an overview of how to convert a project from a monolith to a microservices architecture.

## Code Generation

One of the major goals of this project is to show how to quickly build integrations. Part of that is utilizing code generation to lower the amount of code that needs to be written and maintained. More information can be found in the [Code Generation](./docs/code-generation.md) document.

## Testing

[Testing](./docs/testing.md) contains an overview of how to test the various components of the application.

## CI/CD

Github Actions can be found in the [.github/workflows](./.github/workflows) directory. You can run them locally using [act](https://github.com/nektos/act).
