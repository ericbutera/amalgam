package convert

import (
	"context"

	"github.com/99designs/gqlgen/graphql"
	pb "github.com/ericbutera/amalgam/pkg/feeds/v1"
	"github.com/vektah/gqlparser/gqlerror"
	"google.golang.org/grpc/status"
)

func ValidationToGraphErr(ctx context.Context, s *status.Status) error {
	for _, detail := range s.Details() {
		if v, ok := detail.(*pb.ValidationErrors); ok {
			for _, err := range v.Errors {
				graphql.AddError(ctx, gqlerror.Errorf(err.Message))
			}
			return gqlerror.Errorf("validation error")
		}
	}
	return nil
}
