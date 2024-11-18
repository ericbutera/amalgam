package interceptors

import (
	"context"

	pb "github.com/ericbutera/amalgam/pkg/feeds/v1"
	"github.com/ericbutera/amalgam/services/rpc/internal/server/observability"
	"google.golang.org/grpc"
)

func UnaryMetricMiddlewareHandler(feedMetrics *observability.FeedMetrics) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any, err error) {
		resp, err = handler(ctx, req)
		if err != nil {
			return nil, err
		}

		ProcessMetrics(info.FullMethod, feedMetrics)
		return resp, err
	}
}

func StreamMetricMiddlewareHandler(feedMetrics *observability.FeedMetrics) grpc.StreamServerInterceptor {
	return func(srv any, stream grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) (err error) {
		err = handler(srv, stream)
		if err != nil {
			return err
		}

		ProcessMetrics(info.FullMethod, feedMetrics)
		return err
	}
}

func ProcessMetrics(fullMethod string, metrics *observability.FeedMetrics) {
	switch fullMethod {
	case pb.FeedService_CreateFeed_FullMethodName:
		metrics.FeedsCreated.Inc()
	case pb.FeedService_SaveArticle_FullMethodName:
		metrics.ArticlesCreated.Inc()
	}
}
