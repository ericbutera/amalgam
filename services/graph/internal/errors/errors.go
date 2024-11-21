package errors

import (
	"context"

	"github.com/ericbutera/amalgam/services/graph/internal/convert"
	"github.com/samber/lo"
	"github.com/vektah/gqlparser/gqlerror"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	ErrNotFound      = gqlerror.Errorf("record not found")
	ErrAlreadyExists = gqlerror.Errorf("record already exists")
)

func HandleGrpcErrors(ctx context.Context, err error, msg string) error {
	s := status.Convert(err)
	switch s.Code() { //nolint:exhaustive
	case codes.NotFound:
		return ErrNotFound
	case codes.AlreadyExists:
		return ErrAlreadyExists
	case codes.InvalidArgument:
		return convert.ValidationToGraphErr(ctx, s)
	}
	return gqlerror.Errorf(lo.CoalesceOrEmpty(msg, "could not perform action")) //nolint:govet
}
