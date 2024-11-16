# Development Guide

## Prerequisites

- [docker desktop](https://docs.docker.com/desktop/) + [docker kubernetes](https://docs.docker.com/desktop/features/kubernetes/)
- [asdf-vm](https://asdf-vm.com/) - easily manage third party tools; this is not a hard requirement, but all of the commands things in [.tool-versions](../.tool-versions) are used throughout local development.
- [tilt](https://tilt.dev/) - local development orchestrator

## Tilt

Tilt is the local development orchestrator. The [Tiltfile](../Tiltfile) configures services, dependencies and build steps.

## Local Debugging

I like to debug locally using by turning off specific services in Tilt and running them locally in VSCode.

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