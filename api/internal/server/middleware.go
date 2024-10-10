package server

import (
	"log/slog"

	sloggin "github.com/samber/slog-gin"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
)

func (s *server) middleware() {
	// https://github.com/open-telemetry/opentelemetry-go-contrib/blob/main/instrumentation/github.com/gin-gonic/gin/otelgin/example/server.go
	s.router.Use(otelgin.Middleware(MiddlewareName))
	s.router.Use(sloggin.New(slog.Default()))
}
