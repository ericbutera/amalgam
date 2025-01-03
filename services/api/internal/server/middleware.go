package server

import (
	"log/slog"
	"slices"
	"time"

	"github.com/ericbutera/amalgam/services/api/internal/config"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	sloggin "github.com/samber/slog-gin"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
)

const MaxAge = 12 * time.Hour

var ignoredRoutes = []string{
	"/metrics",
	"/health",
	"/swagger/*any",
}

func (s *server) middleware() {
	s.router.Use(gin.Recovery())
	s.router.Use(logMiddleware())
	s.router.Use(corsMiddleware(s.config))
	if s.config.OtelEnable {
		s.router.Use(otelMiddleware())
	}
}

func otelMiddleware() gin.HandlerFunc {
	// otel: https://github.com/open-telemetry/opentelemetry-go-contrib/blob/main/instrumentation/github.com/gin-gonic/gin/otelgin/example/server.go
	return otelgin.Middleware(
		MiddlewareName,
		otelgin.WithGinFilter(func(c *gin.Context) bool {
			return !slices.Contains(ignoredRoutes, c.Request.URL.Path)
		}),
	)
}

func corsMiddleware(cfg *config.Config) gin.HandlerFunc {
	return cors.New(cors.Config{
		AllowOrigins:     cfg.CorsAllowOrigins,
		AllowMethods:     cfg.CorsAllowMethods,
		AllowHeaders:     cfg.CorsAllowHeaders,
		ExposeHeaders:    cfg.CorsExposeHeaders,
		AllowCredentials: true,
		MaxAge:           MaxAge,
	})
}

func logMiddleware() gin.HandlerFunc {
	return sloggin.NewWithFilters(
		slog.Default(),
		sloggin.IgnorePath(ignoredRoutes...),
	)
}
