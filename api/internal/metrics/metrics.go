package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	TestCounter = promauto.NewCounter(prometheus.CounterOpts{
		Name: "api_test_counter",
		Help: "test counter!",
	})
)
