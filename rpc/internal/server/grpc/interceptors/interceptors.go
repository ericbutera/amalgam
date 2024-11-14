package interceptors

import (
	"context"

	"github.com/ericbutera/amalgam/rpc/internal/server/observability"
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
	case "/feeds.v1.FeedService/CreateFeed":
		metrics.FeedsCreated.Inc()
	case "/feeds.v1.FeedService/CreateArticle":
		metrics.ArticlesCreated.Inc()
	}
}
