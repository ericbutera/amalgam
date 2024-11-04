package client

import (
	"log"
	"time"

	"github.com/ericbutera/amalgam/internal/logger"
	"go.temporal.io/sdk/client"

	prom "github.com/prometheus/client_golang/prometheus"
	"github.com/uber-go/tally/v4"
	"github.com/uber-go/tally/v4/prometheus"
	sdktally "go.temporal.io/sdk/contrib/tally"
)

func NewTemporalClient(host string) (client.Client, error) {
	opts := client.Options{
		Logger:   logger.New(),
		HostPort: host,
		MetricsHandler: sdktally.NewMetricsHandler(newPrometheusScope(prometheus.Configuration{
			ListenAddress: "0.0.0.0:9090",
			TimerType:     "histogram",
		})),
	}
	c, err := client.Dial(opts)
	if err != nil {
		return nil, err
	}
	return c, nil
}

func newPrometheusScope(c prometheus.Configuration) tally.Scope {
	// source: https://github.com/temporalio/samples-go/blob/main/metrics/worker/main.go
	reporter, err := c.NewReporter(
		prometheus.ConfigurationOptions{
			Registry: prom.NewRegistry(),
			OnError: func(err error) {
				log.Println("error in prometheus reporter", err)
			},
		},
	)
	if err != nil {
		log.Fatalln("error creating prometheus reporter", err)
	}
	scopeOpts := tally.ScopeOptions{
		CachedReporter:  reporter,
		Separator:       prometheus.DefaultSeparator,
		SanitizeOptions: &sdktally.PrometheusSanitizeOptions,
		Prefix:          "feed_",
	}
	scope, _ := tally.NewRootScope(scopeOpts, time.Second)
	scope = sdktally.NewPrometheusNamingScope(scope)

	return scope
}
