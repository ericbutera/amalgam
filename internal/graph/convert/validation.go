package convert

import (
	"context"

	"github.com/99designs/gqlgen/graphql"
	"github.com/vektah/gqlparser/gqlerror"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/status"
)

var (
	ErrValidation     = gqlerror.Error{Message: "validation error"}
	ErrInvalidRequest = gqlerror.Error{Message: "invalid request"}
)

func ValidationToGraphErr(ctx context.Context, s *status.Status) error {
	counter := 0

	for _, detail := range s.Details() {
		if br, ok := detail.(*errdetails.BadRequest); ok {
			for _, violation := range br.GetFieldViolations() {
				graphql.AddError(ctx, &gqlerror.Error{
					Message: violation.GetDescription(),
					Extensions: map[string]any{
						"field": violation.GetField(),
					},
				})

				counter++
			}
		}
	}

	if counter > 0 {
		return &ErrValidation
	}

	return &ErrInvalidRequest
}
