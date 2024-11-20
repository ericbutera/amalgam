// https://github.com/grafana/docker-otel-lgtm/blob/main/examples/go/otel.go
package otel

// TODO: i can't seem to find metric & log exporter data in LGTM
// TODO: sample traces (with keep parent)
// TODO: tail call sampling

import (
	"context"
	"errors"
	"os"
	"time"

	"github.com/ericbutera/amalgam/pkg/otel/samplers"
	"go.opentelemetry.io/contrib/bridges/otelslog"
	"go.opentelemetry.io/contrib/instrumentation/runtime"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlplog/otlploghttp"
	"go.opentelemetry.io/otel/exporters/otlp/otlpmetric/otlpmetrichttp"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
	"go.opentelemetry.io/otel/log/global"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/log"
	"go.opentelemetry.io/otel/sdk/metric"
	"go.opentelemetry.io/otel/sdk/trace"
)

const (
	schemaName = "https://github.com/grafana/docker-otel-lgtm"
)

var (
	Tracer = otel.Tracer(schemaName)
	Logger = otelslog.NewLogger(schemaName)
)

// bootstrap the OpenTelemetry pipeline
// If it does not return an error, make sure to call shutdown for proper cleanup.
func Setup(ctx context.Context, ignoredSpans []string) (shutdown func(context.Context) error, err error) {
	var shutdownFuncs []func(context.Context) error

	Logger.Info("setting up OpenTelemetry")

	shutdown = createShutdownFunc(&shutdownFuncs)

	// TODO: remove, if setup is called it's enabled
	if os.Getenv("OTEL_ENABLE") != "true" {
		Logger.Info("OpenTelemetry is disabled")
		return nil, nil
	}

	handleErr := func(inErr error) {
		err = errors.Join(inErr, shutdown(ctx))
	}

	setupPropagators()

	if err := setupTracing(ctx, ignoredSpans, &shutdownFuncs); err != nil {
		handleErr(err)
		return nil, err
	}

	if err := setupMetrics(ctx, &shutdownFuncs); err != nil {
		handleErr(err)
		return nil, err
	}

	if err := setupLogging(ctx, &shutdownFuncs); err != nil {
		handleErr(err)
		return nil, err
	}

	if err := setupRuntimeInstrumentation(); err != nil {
		handleErr(err)
		return nil, err
	}

	return shutdown, nil
}

func createShutdownFunc(shutdownFuncs *[]func(context.Context) error) func(context.Context) error {
	return func(ctx context.Context) error {
		var err error
		for _, fn := range *shutdownFuncs {
			err = errors.Join(err, fn(ctx))
		}
		*shutdownFuncs = nil
		return err
	}
}

func setupPropagators() {
	prop := propagation.NewCompositeTextMapPropagator(
		propagation.TraceContext{},
		propagation.Baggage{},
	)
	otel.SetTextMapPropagator(prop)
}

func setupTracing(ctx context.Context, ignoredSpans []string, shutdownFuncs *[]func(context.Context) error) error {
	exporter, err := otlptrace.New(ctx, otlptracehttp.NewClient())
	if err != nil {
		return err
	}

	opts := []trace.TracerProviderOption{
		trace.WithBatcher(exporter),
	}

	if ignoredSpans != nil {
		opts = append(opts, trace.WithSampler(samplers.NewSpanName(ignoredSpans)))
	}

	provider := trace.NewTracerProvider(opts...)
	*shutdownFuncs = append(*shutdownFuncs, provider.Shutdown)
	otel.SetTracerProvider(provider)
	return nil
}

func setupMetrics(ctx context.Context, shutdownFuncs *[]func(context.Context) error) error {
	exporter, err := otlpmetrichttp.New(ctx)
	if err != nil {
		return err
	}

	provider := metric.NewMeterProvider(metric.WithReader(metric.NewPeriodicReader(exporter)))
	*shutdownFuncs = append(*shutdownFuncs, provider.Shutdown)
	otel.SetMeterProvider(provider)
	return nil
}

func setupLogging(ctx context.Context, shutdownFuncs *[]func(context.Context) error) error {
	exporter, err := otlploghttp.New(ctx, otlploghttp.WithInsecure())
	if err != nil {
		return err
	}

	provider := log.NewLoggerProvider(log.WithProcessor(log.NewBatchProcessor(exporter)))
	*shutdownFuncs = append(*shutdownFuncs, provider.Shutdown)
	global.SetLoggerProvider(provider)
	return nil
}

func setupRuntimeInstrumentation() error {
	return runtime.Start(runtime.WithMinimumReadMemStatsInterval(time.Second))
}
