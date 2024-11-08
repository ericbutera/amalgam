package observability

import (
	grpcprom "github.com/grpc-ecosystem/go-grpc-middleware/providers/prometheus"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/collectors"
)

type Observability struct {
	Registry      *prometheus.Registry
	ServerMetrics *grpcprom.ServerMetrics
	// panicsTotal   prometheus.Counter
}

func New() *Observability {
	registry := prometheus.NewRegistry()
	srvMetrics := newServerMetrics()

	registry.MustRegister(srvMetrics)
	registry.MustRegister(collectors.NewGoCollector(collectors.WithGoCollectorRuntimeMetrics()))

	// TODO: server.newPromMetrics(registry)

	return &Observability{
		Registry:      registry,
		ServerMetrics: srvMetrics,
	}
}

func newServerMetrics() *grpcprom.ServerMetrics {
	return grpcprom.NewServerMetrics(
		grpcprom.WithServerHandlingTimeHistogram(
			grpcprom.WithHistogramBuckets([]float64{0.001, 0.01, 0.1, 0.3, 0.6, 1, 3, 6, 9, 20, 30, 60, 90, 120}),
		),
	)
}

// TODO:
// func (s *Server) newPromMetrics(reg prometheus.Registerer) {
// 	s.panicsTotal = promauto.With(reg).NewCounter(prometheus.CounterOpts{
// 		Name: "grpc_req_panics_recovered_total",
// 		Help: "Total number of gRPC requests recovered from internal panic.",
// 	})
// }
