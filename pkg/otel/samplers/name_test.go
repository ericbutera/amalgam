package samplers_test

import (
	"testing"

	"github.com/ericbutera/amalgam/pkg/otel/samplers"
	"github.com/stretchr/testify/assert"
	"go.opentelemetry.io/otel/sdk/trace"
)

func TestSpanNameSampler(t *testing.T) {
	ignoredNames := []string{"ignored-span-1", "ignored-span-2"}
	sampler := samplers.NewSpanName(ignoredNames)

	tests := []struct {
		spanName         string
		expectedDecision trace.SamplingDecision
	}{
		{"ignored-span-1", trace.Drop},
		{"ignored-span-2", trace.Drop},
		{"normal-span", trace.RecordAndSample},
	}

	for _, tt := range tests {
		t.Run(tt.spanName, func(t *testing.T) {
			params := trace.SamplingParameters{
				Name: tt.spanName,
			}
			result := sampler.ShouldSample(params)
			assert.Equal(t, tt.expectedDecision, result.Decision)
		})
	}
}
