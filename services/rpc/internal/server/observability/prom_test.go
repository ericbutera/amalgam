package observability_test

import (
	"testing"

	"github.com/ericbutera/amalgam/services/rpc/internal/server/observability"
	"github.com/prometheus/client_golang/prometheus/testutil"
	"github.com/stretchr/testify/assert"
)

func TestFeedMetrics(t *testing.T) {
	t.Parallel()

	o := observability.New()
	feedMetrics := o.FeedMetrics
	assert.NotNil(t, feedMetrics.ArticlesCreated)
	assert.NotNil(t, feedMetrics.FeedsCreated)
	assert.NotNil(t, feedMetrics.PanicsTotal)
}

func TestFeedMetrics_ArticlesCreated(t *testing.T) {
	t.Parallel()

	o := observability.New()
	feedMetrics := o.FeedMetrics

	initialValue := testutil.ToFloat64(feedMetrics.ArticlesCreated)
	feedMetrics.ArticlesCreated.Inc()
	assert.InDelta(t, initialValue+1, testutil.ToFloat64(feedMetrics.ArticlesCreated), 0.0001)
}
