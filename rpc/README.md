# RPC

This contains a gRPC service which shows how to extract functionality from a classic rest API and expose it as a gRPC service.

## Validation

[Validation](https://github.com/grpc-ecosystem/go-grpc-middleware) is available as a convenience to consumers of this API. Actual validation is done in the service layer. The earlier validation can happen, the better it will be for system health. However, this is a balancing act as each layer can create inconsistencies.
