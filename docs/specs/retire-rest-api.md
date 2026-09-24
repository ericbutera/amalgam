# Retire the REST API

## Objective

Make GraphQL the only public application API, keep gRPC available only for
internal service communication, and retire the legacy REST service in
`services/api`.

The REST service currently forwards requests to GraphQL. It is already marked
deprecated in [docs/services.md](../services.md), but it is still built,
deployed, monitored, and used by the generated OpenAPI k6 fixture.

## Implementation status

The in-repository retirement is implemented. Production rollout remains gated
by the external-consumer decision below, which cannot be proven from this
repository alone.

## Decision gate: external REST consumers

Do not delete `services/api` until this gate is complete. Repository searches
cannot prove that an external integration is absent.

- [ ] Inventory production ingress, DNS, API gateway routes, dashboards, and
      access logs for REST endpoints and `/swagger`.
- [ ] Contact or inspect every known third-party integration and confirm
      whether it calls the REST API, depends on its OpenAPI document, or can
      move to GraphQL.
- [ ] Search deployment configuration, secrets, runbooks, and customer-facing
      documentation outside this repository for `api`, `/v1`, and REST URLs.
- [ ] Record the result and owner in this document or a linked migration issue.
- [ ] If any supported consumer remains, stop this retirement plan and choose
      one of these paths:
      - [ ] Keep a versioned REST compatibility adapter with an owner,
            authentication, rate limits, monitoring, and a sunset date.
      - [ ] Migrate the consumer to GraphQL, verify parity, and obtain sign-off
            before continuing.
- [ ] Obtain an explicit decision that no supported REST consumer blocks
      deletion.

## 1. Establish GraphQL as the public contract

- [ ] Compare every REST operation with the GraphQL schema and resolver:
      `GET /feeds`, `POST /feeds`, `PUT /feeds/:id`, `GET /feeds/:id`,
      `GET /feeds/:id/articles`, `GET /articles/:id`, and `/health`.
- [ ] Confirm GraphQL has equivalent behavior for validation, errors,
      pagination, authorization, and feed/article identifiers.
- [x] Confirm the UI and `amalgam-cli` use generated GraphQL clients and do not
      depend on REST or OpenAPI-generated clients.
- [ ] Confirm GraphQL owns public concerns currently expected at the edge:
      authentication, CORS, rate limiting, request logging, metrics, tracing,
      and health/readiness endpoints.
- [ ] Document the public GraphQL endpoint, schema publication process, and
      supported client-generation commands.
- [ ] Add or update GraphQL integration tests for each behavior previously
      covered only through REST.

## 2. Remove the REST runtime and deployment path

- [x] Delete `services/api` from the repository implementation after the
      in-repository GraphQL parity and caller checks.
- [x] Delete `helm/templates/api.yaml` and remove the API image, service,
      deployment, probes, environment variables, and observability labels.
- [x] Remove API compilation, image, port-forward, Swagger link, and resource
      dependencies from `Tiltfile` and `Tiltfile.tests`.
- [x] Remove REST-specific API configuration and secrets from Helm values,
      Compose/Tilt configuration, scripts, and CI workflows.
- [x] Remove the `api:8080` target from
      `containers/lgtm/prometheus.yaml` and replace it with GraphQL or another
      intentionally monitored public/internal endpoint.
- [ ] Remove REST-specific ingress, service discovery, network policy, and
      alerting references wherever they exist.
- [ ] Verify the deployed graph is the only public application endpoint and
      the RPC service remains reachable only from approved internal workloads
      or local development port-forwards.

## 3. Retire OpenAPI and REST-generated artifacts

OpenAPI is an artifact of the REST contract. It should not remain an active
contract after the REST server is removed.

- [x] Remove the `generate-openapi`, `generate-api-clients`,
      `generate-go-api-client`, and `generate-k6` REST portions from
      `mise.toml`.
- [x] Remove the `swag` tool and Swagger/OpenAPI-only Go dependencies from the
      project toolchain and `go.mod`/`go.sum` after confirming no other package
      uses them.
- [x] Delete `services/api/docs`, Swagger UI helpers, and REST-generated client
      output that has no remaining consumer.
- [x] Delete or replace `k6/tests/openapi` and its Docker/task wiring. Prefer a
      GraphQL k6 test that exercises the public endpoint and checks the same
      feed/article workflows.
- [x] Remove `API_HOST` and REST assumptions from `helm/templates/k6.yaml`,
      k6 Dockerfiles, generated scripts, and test documentation.
- [ ] If the historical OpenAPI document must be retained for archival or
      migration purposes, move it to an explicitly archived location and label
      it as unsupported; do not continue generating clients or tests from it.

## 4. Update documentation and ownership

- [x] Update `docs/services.md` to remove the deprecated REST service section
      and state the final GraphQL-public/gRPC-internal boundary.
- [x] Update `docs/architecture.md` so the final architecture is clear and
      historical REST evolution is separated from the current design.
- [x] Update `docs/code-generation.md` and `k6/tests/README.md` to describe
      GraphQL schema/client generation and GraphQL load testing.
- [x] Update development, testing, runbook, and onboarding documentation that
      references port 8080 as the REST API or links to Swagger.
- [x] Remove stale repository links to the old `api/` path and REST client
      directories.
- [ ] Record GraphQL schema ownership, gRPC proto ownership, compatibility
      policy, and the process for deprecating public GraphQL fields.

## 5. Dependency and source cleanup

- [x] Remove unused Gin, Swagger, REST client, and API-only dependencies after
      running Go module analysis.
- [x] Remove API-only converters, DTOs, configuration, and tests only after
      verifying they are not shared by GraphQL or gRPC.
- [x] Remove API-specific generated files and update generated-code manifests
      or inventories if present.
- [x] Run a repository-wide search for `services/api`, `api:8080`, `API_HOST`,
      `/v1`, `swagger`, `openapi`, and `api-image`; classify every remaining
      match as intentional history, GraphQL/RPC usage, or stale reference.
- [ ] Do not remove shared domain/service logic that is still owned by GraphQL,
      gRPC, workers, or the data pipeline.

## 6. Validation and release gate

- [ ] Regenerate GraphQL server code, schema, and Go/TypeScript clients; ensure
      generated output is clean.
- [x] Run the equivalent of `mise run go-test` directly with a writable module
      cache: `go test -short -timeout 30s ./...`.
- [ ] Run `mise run go-lint` and `mise run ts-lint`.
- [ ] Run `mise run buf-lint`.
- [ ] Render and validate Helm/Tilt configuration; confirm no API resource or
      `api:8080` dependency remains.
- [ ] Run GraphQL integration tests and the GraphQL k6 smoke/load test against
      the local deployment.
- [ ] Verify the CLI, UI, workers, and data pipeline against GraphQL/gRPC after
      the API removal.
- [ ] Verify observability dashboards, alerts, probes, and logs contain no
      expected REST API signals.
- [ ] Deploy to a non-production environment and confirm GraphQL public access
      and internal-only gRPC access before production rollout.
- [ ] Roll out the removal, monitor GraphQL errors and traffic, and retain a
      rollback plan until the observation window closes.
- [ ] Mark this specification complete only after the external-consumer gate,
      runtime removal, documentation cleanup, and all validation checks pass.

## Completion criteria

- [ ] GraphQL is the documented and deployed public application API.
- [ ] gRPC is used for internal service communication and is not publicly
      routed.
- [ ] No REST API deployment, build target, OpenAPI generation path, or active
      REST client/test remains.
- [ ] No supported external integration depends on the retired REST surface,
      or an explicitly owned compatibility adapter remains instead.
- [ ] The repository search and validation checks are clean.
