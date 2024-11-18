//nolint:err113
package convert

import (
	"context"
	"errors"
	"strings"

	"github.com/99designs/gqlgen/graphql"
	pb "github.com/ericbutera/amalgam/pkg/feeds/v1"
	"github.com/vektah/gqlparser/gqlerror"
	"google.golang.org/grpc/status"
)

var (
	ErrValidation     = gqlerror.Error{Message: "validation error"}
	ErrInvalidRequest = gqlerror.Error{Message: "invalid request"}
)

func ValidationToGraphErr(ctx context.Context, s *status.Status) error {
	for _, detail := range s.Details() {
		if v, ok := detail.(*pb.ValidationErrors); ok {
			for _, err := range v.GetErrors() {
				graphql.AddError(ctx, errors.New(err.GetMessage()))
			}
			return &ErrValidation
		}
	}
	if strings.Contains(s.Message(), "validation") {
		// TODO: modify middleware to use validation errors proto
		// these errors are not friendly
		return errors.New(s.Message())
	}
	return &ErrInvalidRequest
}
