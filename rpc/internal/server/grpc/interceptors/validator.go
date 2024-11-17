package interceptors

import (
	"context"

	"github.com/bufbuild/protovalidate-go"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
)

type Validator struct {
	Unary     grpc.UnaryServerInterceptor
	Stream    grpc.StreamServerInterceptor
	Validator *protovalidate.Validator
}

func NewValidator() Validator {
	v, err := protovalidate.New()
	if err != nil {
		panic(err)
	}

	return Validator{
		Unary:  UnaryValidation(v),
		Stream: StreamValidation(v),
	}
}

func UnaryValidation(validator *protovalidate.Validator) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any, err error) {
		if err := validateRequest(validator, req); err != nil {
			return nil, status.Error(codes.InvalidArgument, err.Error())
		}

		resp, err = handler(ctx, req)
		if err != nil {
			return nil, err
		}

		return resp, err
	}
}

func StreamValidation(validator *protovalidate.Validator) grpc.StreamServerInterceptor {
	return func(srv any, stream grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) (err error) {
		var req any
		if err := stream.RecvMsg(&req); err != nil {
			return err
		}

		if err := validateRequest(validator, req); err != nil {
			return status.Error(codes.InvalidArgument, err.Error())
		}

		err = handler(srv, stream)
		if err != nil {
			return err
		}

		return err
	}
}

func validateRequest(validator *protovalidate.Validator, req any) error {
	msg, ok := req.(proto.Message)
	if !ok {
		return status.Errorf(codes.InvalidArgument, "invalid request")
	}
	if err := validator.Validate(msg); err != nil {
		return status.Error(codes.InvalidArgument, err.Error())
	}
	return nil
}
