package middleware

import (
	"context"
	"log/slog"

	"github.com/99designs/gqlgen/graphql"
)

type ErrorLogging struct{}

func NewErrorLogging() ErrorLogging {
	return ErrorLogging{}
}

func (ErrorLogging) ExtensionName() string {
	return "ErrorLoggingMiddleware"
}

func (ErrorLogging) Validate(_ graphql.ExecutableSchema) error {
	return nil
}

func (ErrorLogging) InterceptField(ctx context.Context, next graphql.Resolver) (res any, err error) {
	res, err = next(ctx)
	if err != nil {
		slog.Error("graph error", "error", err)
	}

	return res, err
}
