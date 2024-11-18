# Development Guide

## Prerequisites

- [docker desktop](https://docs.docker.com/desktop/) + [docker kubernetes](https://docs.docker.com/desktop/features/kubernetes/)
- [asdf-vm](https://asdf-vm.com/) - easily manage third party tools (or manually install things listed in [.tool-versions](../.tool-versions)).
- [tilt](https://tilt.dev/) - local development orchestrator (installed via asdf).

## Tilt

Tilt is the local development orchestrator. The [Tiltfile](../Tiltfile) configures services, dependencies and build steps.

## Linters

Be sure to install the pre-commit hooks which run various linters, formatters, and tests.

```sh
just setup
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
      "program": "graph/main.go",
      "envFile": "${workspaceFolder}/.env",
      "args": ["server"]
    },
    {
      "name": "rpc",
      "type": "go",
      "request": "launch",
      "mode": "debug",
      "program": "rpc/main.go",
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
      "program": "data-pipeline/temporal/feed/worker/main.go",
      "envFile": "${workspaceFolder}/.env"
    },
    {
      "name": "temporal - generate feed - worker",
      "type": "go",
      "request": "launch",
      "mode": "debug",
      "program": "data-pipeline/temporal/generate/worker/main.go",
      "envFile": "${workspaceFolder}/.env"
    },
    {
      "name": "ui - firefox",
      "type": "firefox",
      "request": "launch",
      "url": "http://localhost:3000",
      "reAttach": true,
      "pathMappings": [
        {
          "url": "webpack://_n_e/app",
          "path": "${workspaceFolder}/ui/app"
        }
      ]
    }
  ]
}
```

example `.vscode/settings.json`:

```json
{
  "go.testEnvFile": "${workspaceFolder}/.env"
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
