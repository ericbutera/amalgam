package client

import (
	"log"
	"time"

	"github.com/ericbutera/amalgam/internal/logger"
	"github.com/ericbutera/amalgam/pkg/config/env"
	prom "github.com/prometheus/client_golang/prometheus"
	"github.com/uber-go/tally/v4"
	"github.com/uber-go/tally/v4/prometheus"
	"go.temporal.io/sdk/client"
	sdktally "go.temporal.io/sdk/contrib/tally"
)

func NewTemporalClient(host string) (client.Client, error) {
	scope, err := newPrometheusScope(prometheus.Configuration{
		ListenAddress: "0.0.0.0:9090",
		TimerType:     "histogram",
	})
	if err != nil {
		return nil, err
	}

	opts := client.Options{
		Logger:         logger.New(),
		HostPort:       host,
		MetricsHandler: sdktally.NewMetricsHandler(scope),
	}
	c, err := client.Dial(opts)
	if err != nil {
		return nil, err
	}
	return c, nil
}

type Config struct {
	TemporalHost string `mapstructure:"temporal_host"`
}

func NewTemporalClientFromEnv() (client.Client, error) {
	config, err := env.New[Config]()
	if err != nil {
		return nil, err
	}
	c, err := NewTemporalClient(config.TemporalHost)
	if err != nil {
		return nil, err
	}
	return c, nil
}

func newPrometheusScope(c prometheus.Configuration) (tally.Scope, error) {
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
		return nil, err
	}
	scopeOpts := tally.ScopeOptions{
		CachedReporter:  reporter,
		Separator:       prometheus.DefaultSeparator,
		SanitizeOptions: &sdktally.PrometheusSanitizeOptions,
	}
	scope, _ := tally.NewRootScope(scopeOpts, time.Second)
	scope = sdktally.NewPrometheusNamingScope(scope)

	return scope, nil
}
