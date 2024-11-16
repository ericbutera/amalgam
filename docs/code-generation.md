# Code Generation

## Copygen

Copygen is a project that can generate code to copy fields from one struct to another. This is a very common task when using the best practice of separating interface from implementation. For example, GraphQL should not be showing the literal fields that flow thru gRPC to Service Layer to the database. Instead each boundary should have separate data types. In the real world, the fields are often the same, so it makes sense to use a tool like copygen as much as possible.

You can find the [copygen transforms here](https://github.com/ericbutera/amalgam/blob/ad3d79839030889826a8fb2f0c0dcad48bf9d06e/internal/copygen/setup.go#L12-L41).

## GraphQL Clients (post [v1.4.0](https://github.com/ericbutera/amalgam/releases/tag/v1.4.0))

- [TypeScript](https://github.com/ericbutera/amalgam/blob/9528beb51c6b2affa3b6bd1622ca666983148fc4/ui/app/generated/graphql.ts#L204-L225)
- [Go](https://github.com/ericbutera/amalgam/blob/9528beb51c6b2affa3b6bd1622ca666983148fc4/pkg/clients/graphql/graphql.gen.go)

## OpenAPI Clients ([v1.3.1](https://github.com/ericbutera/amalgam/releases/tag/v1.3.1))

- [OpenAPI spec](./api/docs/swagger.yaml) with [swaggo/swag](https://github.com/swaggo/swag)
- [REST client](./pkg/client/README.md) from OpenAPI spec
- [TypeScript client](./ui/app/lib/client/) from OpenAPI spec
- [k6 tests](./k6/README.md) from OpenAPI spec
