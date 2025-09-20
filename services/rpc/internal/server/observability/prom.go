package observability

import (
	grpcprom "github.com/grpc-ecosystem/go-grpc-middleware/providers/prometheus"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/collectors"
)

type Observability struct {
	Registry      *prometheus.Registry
	ServerMetrics *grpcprom.ServerMetrics
	FeedMetrics   *FeedMetrics
}

func New() *Observability {
	registry := prometheus.NewRegistry()
	registry.MustRegister(collectors.NewGoCollector(collectors.WithGoCollectorRuntimeMetrics()))

	o := &Observability{
		Registry:      registry,
		ServerMetrics: newServerMetrics(registry),
		FeedMetrics:   getFeedMetrics(registry),
	}

	return o
}

func newServerMetrics(r *prometheus.Registry) *grpcprom.ServerMetrics {
	m := grpcprom.NewServerMetrics(
		grpcprom.WithServerHandlingTimeHistogram(
			grpcprom.WithHistogramBuckets([]float64{0.001, 0.01, 0.1, 0.3, 0.6, 1, 3, 6, 9, 20, 30, 60, 90, 120}),
		),
	)
	r.MustRegister(m)

	return m
}

type FeedMetrics struct {
	ArticlesCreated prometheus.Counter
	FeedsCreated    prometheus.Counter
	PanicsTotal     prometheus.Counter
}

func getFeedMetrics(r *prometheus.Registry) *FeedMetrics {
	// TODO: move this into New() and use Options pattern?
	// TODO: refactor- seems to be a candidate for a metrics package
	// TODO: if register is called multiple times it will fail

	// According to [Practical Monitoring](https://learning.oreilly.com/library/view/practical-monitoring/9781491957349/ch05.html#Monitoring_the_Business)
	// business KPIs can help guide metric creation. this app is an aggregator,
	// we can see if "things are working" by tracking feed & article creation.
	articlesCreated := prometheus.NewCounter(prometheus.CounterOpts{
		Name:      "articles_created_total",
		Namespace: "rpc",
		Help:      "Total number of articles created.",
	})
	feedsCreated := prometheus.NewCounter(prometheus.CounterOpts{
		Name:      "feeds_created_total",
		Namespace: "rpc",
		Help:      "Total number of feeds created.",
	})
	panicsTotal := prometheus.NewCounter(prometheus.CounterOpts{
		Name:      "panics_recovered_total",
		Namespace: "rpc",
		Help:      "Total number of panics recovered.",
	})

	r.MustRegister(articlesCreated)
	r.MustRegister(feedsCreated)
	r.MustRegister(panicsTotal)

	return &FeedMetrics{
		ArticlesCreated: articlesCreated,
		FeedsCreated:    feedsCreated,
		PanicsTotal:     panicsTotal,
	}
}
