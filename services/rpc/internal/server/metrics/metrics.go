package metrics

import (
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

const Timeout = 10 * time.Second

func NewServer(registry *prometheus.Registry, address string) *http.Server {
	m := http.NewServeMux()
	m.Handle("/metrics", promhttp.HandlerFor(
		registry,
		promhttp.HandlerOpts{
			EnableOpenMetrics: true, // Opt into OpenMetrics e.g. to support exemplars.
		},
	))
	return &http.Server{
		Addr:              address,
		Handler:           m,
		ReadHeaderTimeout: Timeout,
	}
}
