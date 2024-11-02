package app

import (
	"io"
	"log"
	"log/slog"
	"net/http"
	"time"

	"github.com/ericbutera/amalgam/internal/http/transport"
	"github.com/ericbutera/amalgam/internal/logger"
	"go.temporal.io/sdk/client"

	prom "github.com/prometheus/client_golang/prometheus"
	"github.com/uber-go/tally/v4"
	"github.com/uber-go/tally/v4/prometheus"
	sdktally "go.temporal.io/sdk/contrib/tally"
)

const FetchTimeout = 10 * time.Second

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

type FetchCallbackParams struct {
	Reader      io.Reader
	Size        int64
	ContentType string
}

type FetchCallback = func(params FetchCallbackParams) error

func FetchUrl(url string, fetchCb FetchCallback) error {
	// note: do not retry; workflow will handle retries
	// TODO: otel
	client := &http.Client{
		Transport: transport.NewLoggingTransport(
			transport.WithLogger(slog.Default()),
		),
		Timeout: FetchTimeout,
	}
	req, err := http.NewRequest("GET", url, nil)
	req.Header.Set("User-Agent", "curl/7.79.1")
	if err != nil {
		return err
	}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return fetchCb(FetchCallbackParams{
		ContentType: resp.Header.Get("Content-Type"),
		Reader:      resp.Body,
		Size:        resp.ContentLength,
	})
}
