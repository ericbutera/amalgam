package server

import (
	"net/http"

	"github.com/ericbutera/amalgam/api/internal"
	"github.com/ericbutera/amalgam/api/internal/metrics"
	"github.com/gin-gonic/gin"
	"github.com/penglongli/gin-metrics/ginmetrics"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
)

const (
	MiddlewareName = "api"
)

type server struct {
	// Gin Router
	r *gin.Engine
}

func New() *server {
	// TODO: ensure gin uses signal.NotifyContext ctx
	// TODO: ensure gin uses slog as logger (for otel exporter)
	s := &server{r: gin.Default()}
	s.middleware()
	s.metrics()
	s.routes()
	return s
}

func (s *server) Run() error {
	return s.r.Run()
}

func (s *server) middleware() {
	// https://github.com/open-telemetry/opentelemetry-go-contrib/blob/main/instrumentation/github.com/gin-gonic/gin/otelgin/example/server.go
	s.r.Use(otelgin.Middleware(MiddlewareName))
}

func (s *server) metrics() {
	// https://github.com/penglongli/gin-metrics
	m := ginmetrics.GetMonitor()
	m.SetMetricPath("/metrics")
	m.SetSlowTime(10)                              // +optional set slow time, default 5s
	m.SetDuration([]float64{0.1, 0.3, 1.2, 5, 10}) // +optional set request duration, default {0.1, 0.3, 1.2, 5, 10} // used to p95, p99
	m.Use(s.r)
}

func (s *server) routes() {
	s.r.GET("/", func(c *gin.Context) { c.Redirect(http.StatusMovedPermanently, internal.SwaggerUri) })
	s.r.GET("/health", func(c *gin.Context) {
		metrics.TestCounter.Inc()
		c.Status(http.StatusOK)
	})

	// TODO auth
	// auth := r.Group("/")
	// auth.Use(app.AuthRequired()) {
	s.r.GET("/swagger/*any", internal.Swagger())
	//}
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
