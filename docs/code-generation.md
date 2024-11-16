# Code Generation

## Copygen

[Copygen](https://github.com/switchupcb/copygen) is a powerful tool designed to simplify a common development task: transferring data between structs. This is especially useful when adhering to best practices like separating concerns in layered architectures (e.g., GraphQL, gRPC, and data layers).

By automating the generation of code for copying fields between structs with similar structures, Copygen significantly reduces development time and effort. This allows developers to focus on core business logic, while the tool handles the repetitive and error-prone task of manual data mapping."

You can find the [copygen transforms here](https://github.com/ericbutera/amalgam/blob/ad3d79839030889826a8fb2f0c0dcad48bf9d06e/internal/copygen/setup.go#L12-L41).

## GraphQL

General workflow:

- run tilt `tilt up`
- change server schema `graph/graph/schema.graphqls`
- generate server `just generate-graph-server`
- await graph service to hot-reload
- generate schema `just generate-graph-schema`
- generate clients `just generate-graph-clients`

This is not an optimal solution. I intend to have it so these steps are automated without a Tilt dependency.

### Server

The GraphQL server can be generated using `just generate-graph-server` ([source](https://github.com/ericbutera/amalgam/blob/ad3d79839030889826a8fb2f0c0dcad48bf9d06e/justfile#L144-L148)).

### Schema

The GraphQL [schema](https://github.com/ericbutera/amalgam/tree/ad3d79839030889826a8fb2f0c0dcad48bf9d06e/tools/graphql-schema) is generated using `just generate-graph-schema` ([source](https://github.com/ericbutera/amalgam/blob/ad3d79839030889826a8fb2f0c0dcad48bf9d06e/justfile#L151-L156)).

### Clients (post [v1.4.0](https://github.com/ericbutera/amalgam/releases/tag/v1.4.0))

These generated clients provide strongly typed interfaces, enhancing developer experience with code completion and preventing runtime errors.

| Client | Command |
| --- | --- |
| [TypeScript](https://github.com/ericbutera/amalgam/blob/9528beb51c6b2affa3b6bd1622ca666983148fc4/ui/app/generated/graphql.ts#L204-L225) | [`generate-graph-ts-client`](https://github.com/ericbutera/amalgam/blob/ad3d79839030889826a8fb2f0c0dcad48bf9d06e/justfile#L165-L169) |
| [Go](https://github.com/ericbutera/amalgam/blob/9528beb51c6b2affa3b6bd1622ca666983148fc4/pkg/clients/graphql/graphql.gen.go) | [`generate-graph-golang-client`](https://github.com/ericbutera/amalgam/blob/ad3d79839030889826a8fb2f0c0dcad48bf9d06e/justfile#L159-L162) |

## OpenAPI Clients ([v1.3.1](https://github.com/ericbutera/amalgam/releases/tag/v1.3.1))

- [OpenAPI spec](https://github.com/ericbutera/amalgam/blob/8c4e26f23ecd3af6c7eae80cbb1a16165fcd1703/api/docs/swagger.yaml) with [swaggo/swag](https://github.com/swaggo/swag)
- [REST client](https://github.com/ericbutera/amalgam/tree/8c4e26f23ecd3af6c7eae80cbb1a16165fcd1703/pkg/client) from OpenAPI spec
- [TypeScript client](https://github.com/ericbutera/amalgam/tree/8c4e26f23ecd3af6c7eae80cbb1a16165fcd1703/ui/app/lib/client) from OpenAPI spec
- [k6 tests](https://github.com/ericbutera/amalgam/tree/main/k6/tests/openapi) from OpenAPI spec
