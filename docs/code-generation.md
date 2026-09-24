# Code Generation

## Goverter

[Goverter](https://goverter.jmattheis.de/) is a powerful tool designed to simplify a common development task: transferring data between structs. This is especially useful when adhering to best practices like separating concerns in layered architectures (e.g., GraphQL, gRPC, and data layers).

By automating the generation of code for copying fields between structs with similar structures, Goverter significantly reduces development time and effort. This allows developers to focus on core business logic, while the tool handles the repetitive and error-prone task of manual data mapping."

## GraphQL

General workflow:

- run tilt `tilt up`
- change server schema `internal/graph/schema.graphqls`
- generate server `mise run generate-graph-server`
- await graph service to hot-reload
- generate schema `mise run generate-graph-schema`
- generate clients `mise run generate-graph-clients`

This is not an optimal solution. I intend to have it so these steps are automated without a Tilt dependency.

## Clients

### ([v1.4.0](https://github.com/ericbutera/amalgam/releases/tag/v1.4.0)) GraphQL

These generated clients provide strongly typed interfaces, enhancing developer experience with code completion and preventing runtime errors.

| Client                                                                                                                                  | Command                                                                                                                                  |
| --------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------- |
| [TypeScript](https://github.com/ericbutera/amalgam/blob/9528beb51c6b2affa3b6bd1622ca666983148fc4/ui/app/generated/graphql.ts#L204-L225) | `mise run generate-graph-ts-client`     |
| [Go](https://github.com/ericbutera/amalgam/blob/9528beb51c6b2affa3b6bd1622ca666983148fc4/pkg/clients/graphql/graphql.gen.go)            | `mise run generate-graph-golang-client` |

There is no active REST/OpenAPI client-generation workflow. Public clients
should be generated from the GraphQL schema and operations above; internal
service clients are generated from the protobuf definitions with Buf.
