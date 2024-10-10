package server

import "go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"

func (s *server) middleware() {
	// https://github.com/open-telemetry/opentelemetry-go-contrib/blob/main/instrumentation/github.com/gin-gonic/gin/otelgin/example/server.go
	s.router.Use(otelgin.Middleware(MiddlewareName))
}
