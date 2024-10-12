package server

import (
	"log/slog"
	"slices"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	sloggin "github.com/samber/slog-gin"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
)

func (s *server) middleware() {
	ignoredRoutes := []string{"/metrics", "/health"}

	s.router.Use(gin.Recovery())
	s.router.Use(sloggin.NewWithFilters(slog.Default(), sloggin.IgnorePath(ignoredRoutes...)))

	// otel: https://github.com/open-telemetry/opentelemetry-go-contrib/blob/main/instrumentation/github.com/gin-gonic/gin/otelgin/example/server.go
	s.router.Use(otelgin.Middleware(MiddlewareName, otelgin.WithGinFilter(func(c *gin.Context) bool {
		return !slices.Contains(ignoredRoutes, c.Request.URL.Path)
	})))

	s.router.Use(cors.New(cors.Config{
		AllowOrigins:     s.config.CorsAllowOrigins,
		AllowMethods:     s.config.CorsAllowMethods,
		AllowHeaders:     s.config.CorsAllowHeaders,
		ExposeHeaders:    s.config.CorsExposeHeaders,
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
}
