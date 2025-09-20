package samplers

import (
	"log/slog"

	"go.opentelemetry.io/otel/sdk/trace"
)

const DefaultTraceRatio = 1.0 // 1.0 == 100%, not for production

func NewSpanName(ignored []string) trace.Sampler {
	slog.Debug("creating span name sampler", "ignored", ignored)
	s := &SpanName{
		BaseSampler:  trace.TraceIDRatioBased(DefaultTraceRatio), // TODO: configurable
		IgnoredNames: ignored,
	}

	return trace.ParentBased(s)
}

type SpanName struct {
	BaseSampler  trace.Sampler
	IgnoredNames []string
}

func (cs *SpanName) ShouldSample(p trace.SamplingParameters) trace.SamplingResult {
	for _, ignored := range cs.IgnoredNames {
		if p.Name == ignored {
			return trace.SamplingResult{Decision: trace.Drop}
		}
	}

	return cs.BaseSampler.ShouldSample(p)
}

func (*SpanName) Description() string {
	return "Sample by Span Name"
}
