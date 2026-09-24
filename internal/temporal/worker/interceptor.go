package worker

import (
	"context"
	"log/slog"

	"github.com/samber/lo"
	"go.opentelemetry.io/otel/trace"
	"go.temporal.io/sdk/contrib/opentelemetry"
	"go.temporal.io/sdk/interceptor"
)

func NewInterceptors(tracer trace.Tracer) []interceptor.WorkerInterceptor {
	traceInterceptor := lo.Must(opentelemetry.NewTracer(opentelemetry.TracerOptions{
		Tracer: tracer,
	}))

	return []interceptor.WorkerInterceptor{
		interceptor.NewTracingInterceptor(traceInterceptor),
	}
}

func HandleShutdown(ctx context.Context, shutdown func(context.Context) error) {
	slog.Info("shutting down otel")

	err := shutdown(ctx)
	if err != nil {
		slog.Error("failed to shutdown OpenTelemetry", "error", err)
	}
}
