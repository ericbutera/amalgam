package interceptors_test

import (
	"context"
	"testing"

	"github.com/bufbuild/protovalidate-go"
	"github.com/ericbutera/amalgam/internal/test/seed"
	pb "github.com/ericbutera/amalgam/pkg/feeds/v1"
	"github.com/ericbutera/amalgam/services/rpc/internal/server/grpc/interceptors"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
)

func TestUnaryValidation(t *testing.T) {
	validator, err := protovalidate.New()
	require.NoError(t, err)

	interceptor := interceptors.UnaryValidation(validator)
	req := &empty.Empty{}
	ctx := context.Background()
	handler := func(_ context.Context, _ any) (any, error) {
		return &empty.Empty{}, nil
	}

	resp, err := interceptor(ctx, req, nil, handler)
	require.NoError(t, err)
	assert.NotNil(t, resp)
}

type MockValidator struct {
	mock.Mock
	protovalidate.Validator
}

func (m *MockValidator) Validate(msg proto.Message) error {
	args := m.Called(msg)
	return args.Error(0)
}

func TestUnaryValidationError(t *testing.T) {
	validator, err := protovalidate.New()
	require.NoError(t, err)

	interceptor := interceptors.UnaryValidation(validator)
	resp, err := interceptor(
		context.Background(),
		&pb.CreateFeedRequest{
			Feed: &pb.CreateFeedRequest_Feed{
				Name: "a",
				Url:  "invalid-url",
			},
			User: &pb.User{Id: seed.UserID},
		},
		nil, nil)

	require.Error(t, err)
	assert.Nil(t, resp)

	st, ok := status.FromError(err)
	assert.True(t, ok)
	assert.Equal(t, codes.InvalidArgument, st.Code())

	for _, detail := range st.Details() {
		if v, ok := detail.(*errdetails.BadRequest); ok {
			violations := v.GetFieldViolations()
			assert.Len(t, violations, 1)
			assert.Equal(t, "feed.url", violations[0].GetField())

			return
		}
	}

	assert.Fail(t, "validation error not found")
}
