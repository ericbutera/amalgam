package main

import (
	"context"
	"errors"
	"net/http"
	"os"
	"os/signal"

	"github.com/gin-gonic/gin"
	"github.com/penglongli/gin-metrics/ginmetrics"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
)

const MiddlewareName = "api"

func main() {
	// https://github.com/grafana/docker-otel-lgtm/blob/main/examples/go/main.go
	// Handle SIGINT (CTRL+C) gracefully.
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	// Set up OpenTelemetry.
	otelShutdown, err := setupOTelSDK(ctx)
	if err != nil {
		return
	}
	defer func() {
		err = errors.Join(err, otelShutdown(context.Background()))
	}()

	r := gin.Default()

	// https://github.com/open-telemetry/opentelemetry-go-contrib/blob/main/instrumentation/github.com/gin-gonic/gin/otelgin/example/server.go
	r.Use(otelgin.Middleware(MiddlewareName))

	// https://github.com/penglongli/gin-metrics
	m := ginmetrics.GetMonitor()
	m.SetMetricPath("/metrics")
	m.SetSlowTime(10)                              // +optional set slow time, default 5s
	m.SetDuration([]float64{0.1, 0.3, 1.2, 5, 10}) // +optional set request duration, default {0.1, 0.3, 1.2, 5, 10} // used to p95, p99
	m.Use(r)

	r.GET("/health", func(c *gin.Context) {
		c.Status(http.StatusOK)
	})
	srvErr := make(chan error, 1)
	go func() {
		srvErr <- r.Run()
	}()

	select {
	case err = <-srvErr:
		// Error when starting HTTP server.
		return
	case <-ctx.Done():
		// Wait for first CTRL+C.
		stop()
	}

	return
}

/*
Routes:
GET /health
POST /register
POST /login
POST /logout
GET /feeds
POST /feed
	- adds a feed source
	- generic to all users
	- create a user_feed record if exists
GET /feed/:id
GET /feed/:id/articles
GET /article/:id

*/
