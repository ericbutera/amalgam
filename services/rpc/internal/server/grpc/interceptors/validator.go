package interceptors

import (
	"context"
	"errors"

	"github.com/bufbuild/protovalidate-go"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
)

func UnaryValidation(validator *protovalidate.Validator) grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req any,
		_ *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (resp any, err error) {
		switch msg := req.(type) {
		case proto.Message:
			if err = validator.Validate(msg); err != nil {
				st := status.New(codes.InvalidArgument, "validation failed")
				br := &errdetails.BadRequest{}

				for _, violation := range parseValidationErrors(err) {
					br.FieldViolations = append(br.FieldViolations, &errdetails.BadRequest_FieldViolation{
						Field:       violation.Field,
						Description: violation.Description,
					})
				}

				st, _ = st.WithDetails(br)
				return nil, st.Err()
			}
		default:
			return nil, status.New(codes.Internal, "invalid request").Err()
		}

		return handler(ctx, req)
	}
}

type FieldViolation struct {
	Field       string
	Description string
}

func parseValidationErrors(err error) []FieldViolation {
	var violations []FieldViolation

	var validationErr *protovalidate.ValidationError
	if errors.As(err, &validationErr) {
		for _, violation := range validationErr.Violations {
			violations = append(violations, FieldViolation{
				Field:       violation.GetFieldPath(),
				Description: violation.GetMessage(),
			})
		}
	}

	return violations
}
