package interceptors_test

import (
	"context"
	"testing"

	pb "github.com/ericbutera/amalgam/pkg/feeds/v1"
	"github.com/ericbutera/amalgam/rpc/internal/server/grpc/interceptors"
	"github.com/ericbutera/amalgam/rpc/internal/server/observability"
	"github.com/prometheus/client_golang/prometheus/testutil"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
)

func TestUnaryMetricMiddlewareHandler(t *testing.T) {
	o := observability.New()
	handler := interceptors.UnaryMetricMiddlewareHandler(o.FeedMetrics)

	ctx := context.Background()
	req := &pb.CreateFeedRequest{}
	info := &grpc.UnaryServerInfo{
		FullMethod: pb.FeedService_CreateFeed_FullMethodName,
	}
	_, err := handler(ctx, req, info, func(_ context.Context, _ any) (any, error) {
		return nil, nil
	})
	require.NoError(t, err)
	assert.InDelta(t, float64(1.0), testutil.ToFloat64(o.FeedMetrics.FeedsCreated), 0.0001)
}

type mockServerStream struct {
	grpc.ServerStream
}

func TestStreamMetricMiddlewareHandler(t *testing.T) {
	o := observability.New()
	handler := interceptors.StreamMetricMiddlewareHandler(o.FeedMetrics)

	stream := &mockServerStream{}
	info := &grpc.StreamServerInfo{
		FullMethod: pb.FeedService_SaveArticle_FullMethodName,
	}
	err := handler(nil, stream, info, func(_ any, _ grpc.ServerStream) error {
		return nil
	})
	require.NoError(t, err)
	assert.InDelta(t, float64(1.0), testutil.ToFloat64(o.FeedMetrics.ArticlesCreated), 0.0001)
}
