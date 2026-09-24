# Development Guide

## Prerequisites

- [docker desktop](https://docs.docker.com/desktop/) + [docker kubernetes](https://docs.docker.com/desktop/features/kubernetes/)
- [mise](https://mise.jdx.dev/) - manage the pinned tools and project tasks in [mise.toml](../mise.toml).
- [tilt](https://tilt.dev/) - local development orchestrator (installed by mise).

## Tilt

Tilt is the local development orchestrator. The [Tiltfile](../Tiltfile) configures services, dependencies and build steps.

## Go repository layout

The Go code is organized as a tool-neutral monorepo:

- `cmd/` contains executable entrypoints and deployment binaries. Each subdirectory is one `main` package.
- `internal/` contains private application implementations, service wiring, CLI commands, and workflow code.
- `pkg/` contains shared or externally consumable contracts, including generated protobuf and GraphQL client packages.
- `tools/` contains developer and code-generation tooling rather than application services.
- `integration/` contains cross-package and infrastructure-facing tests.

Use `go run ./cmd/<name>` or the corresponding `mise` task to run a binary. Keep reusable code below `internal/` or `pkg/`; do not add new application services under a top-level `services/` directory.

## Linters

Be sure to install the pre-commit hooks which run various linters, formatters, and tests.

```sh
mise run setup
```

A few of the linters used:

- [golangci-lint](https://golangci-lint.run/) - golang linters (config in [.golangci.yml](../.golangci.yaml))
- [eslint](https://eslint.org/) - typescript linters
- [conventional commits](https://www.conventionalcommits.org/en/v1.0.0/) - enforced by pre-commit hooks

Check the [github actions](https://github.com/ericbutera/amalgam/blob/5ab8ab5ed5d12669f7258025cfadcc4f0a968ff6/.github/workflows/all.yaml) and [.pre-commit-config.yaml](https://github.com/ericbutera/amalgam/blob/5ab8ab5ed5d12669f7258025cfadcc4f0a968ff6/.pre-commit-config.yaml) for the full list.

## Local Debugging

I like to debug locally in VScode. For example, to debug RPC, first turn off theRPC service in Tilt. Next, run the rpc launch profile. The `.env` file defines environment variables that match the port forwards in Tilt.

example `.vscode/launch.json`:

```json
{
  "version": "0.2.0",
  "configurations": [
    {
      "name": "graph",
      "type": "go",
      "request": "launch",
      "mode": "debug",
      "program": "cmd/graph",
      "envFile": "${workspaceFolder}/.env",
      "args": ["server"]
    },
    {
      "name": "rpc",
      "type": "go",
      "request": "launch",
      "mode": "debug",
      "program": "cmd/rpc",
      "env": {
        "PORT": "50055",
        "METRIC_ADDRESS": ":9091",
        "DB_ADAPTER": "sqlite"
      },
      "envFile": "${workspaceFolder}/.env",
      "args": ["server"]
    },
    {
      "name": "temporal - fetch feed - worker",
      "type": "go",
      "request": "launch",
      "mode": "debug",
      "program": "cmd/feed-fetch-worker",
      "envFile": "${workspaceFolder}/.env"
    },
    {
      "name": "temporal - generate feed - worker",
      "type": "go",
      "request": "launch",
      "mode": "debug",
      "program": "cmd/feed-tasks-worker",
      "envFile": "${workspaceFolder}/.env"
    },
    {
      "name": "ui: debug client-side (Firefox)",
      "type": "firefox",
      "request": "launch",
      "url": "http://localhost:3000",
      "reAttach": true,
      "pathMappings": [
        {
          "url": "webpack://_n_e",
          "path": "${workspaceFolder}/ui"
        },
        {
          "url": "webpack://_n_e/node_modules/",
          "path": "${workspaceFolder}/ui/node_modules"
        }
      ]
    }
  ]
}
```

example `.vscode/settings.json`:

```json
{
  // go specific
  "go.lintOnSave": "workspace",
  "go.testEnvFile": "${workspaceFolder}/.env"
    "eslint.validate": [
    "javascript",
    "typescript",
    "typescriptreact",
    "javascriptreact"
  ],
  "[go]": {
    "editor.formatOnSave": true,
    "editor.defaultFormatter": "golang.go",
    "editor.codeActionsOnSave": {
      "source.organizeImports": "explicit"
    }
  },

  // typescript specific
  "eslint.format.enable": true,
  "[typescript]": {
    "editor.defaultFormatter": "esbenp.prettier-vscode",
    "editor.tabSize": 2,
    "editor.formatOnPaste": true,
    "editor.formatOnSave": true,
    "editor.codeActionsOnSave": {
      "source.fixAll": "explicit",
      "organizeImports": "explicit"
    }
  },
  "[typescriptreact]": {
    "editor.defaultFormatter": "esbenp.prettier-vscode",
    "editor.tabSize": 2,
    "editor.formatOnPaste": true,
    "editor.formatOnSave": true,
    "editor.codeActionsOnSave": {
      "source.fixAll": "explicit",
      "organizeImports": "explicit"
    }
  }
}
```

example `.env`:

```sh
# FAKE_HOST=localhost:8084
FAKE_HOST=faker:8080
GRAPH_HOST="http://localhost:8082/query"
GRAPH_PORT=8082
RPC_HOST=localhost:50055
RPC_INSECURE=true
TEMPORAL_HOST=localhost:7233
MINIO_ENDPOINT=localhost:9100
MINIO_USE_SSL=false
MINIO_ACCESS_KEY=minio
MINIO_SECRET_ACCESS_KEY=minio-password
OTEL_ENABLE=true
OTEL_SERVICE_NAME=feed-worker
OTEL_EXPORTER_OTLP_INSECURE=true
OTEL_EXPORTER_OTLP_ENDPOINT=http://localhost:4318
USE_SCHEDULE=false
FEED_SCHEDULE_ID=feed-worker
FEED_WORKFLOW_ID=feed-worker
CORS_ALLOW_ORIGINS=http://localhost:3000
CORS_ALLOW_METHODS=GET,POST,PUT
CORS_ALLOW_HEADERS=Content-Type,Authorization,Origin
CORS_EXPOSE_HEADERS=Content-Length
```
