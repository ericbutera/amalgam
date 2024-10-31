package logger

import (
	"log/slog"
	"os"
)

func New() *slog.Logger {
	// TODO: add traceID to logs (this is from grpc ecosystem)
	// logTraceID := func(ctx context.Context) logging.Fields {
	// 	if span := trace.SpanContextFromContext(ctx); span.IsSampled() {
	// 		return logging.Fields{"traceID", span.TraceID().String()}
	// 	}
	// 	return nil
	// }
	// opts := []logging.Option{
	// 	logging.WithLogOnEvents(logging.StartCall, logging.FinishCall),
	// 	logging.WithFieldsFromContext(logTraceID),
	// }

	handler := slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	})
	logger := slog.New(handler)
	slog.SetDefault(logger)
	return logger
}
