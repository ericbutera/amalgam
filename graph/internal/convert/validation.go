//nolint:err113
package convert

import (
	"context"
	"errors"

	"github.com/99designs/gqlgen/graphql"
	pb "github.com/ericbutera/amalgam/pkg/feeds/v1"
	"github.com/vektah/gqlparser/gqlerror"
	"google.golang.org/grpc/status"
)

func ValidationToGraphErr(ctx context.Context, s *status.Status) error {
	for _, detail := range s.Details() {
		if v, ok := detail.(*pb.ValidationErrors); ok {
			for _, err := range v.GetErrors() {
				graphql.AddError(ctx, errors.New(err.GetMessage()))
			}
			return &gqlerror.Error{Message: "validation error"}
		}
	}
	return nil
}
