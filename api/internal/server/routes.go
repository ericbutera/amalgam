package server

import (
	"net/http"

	"github.com/ericbutera/amalgam/api/internal"
	"github.com/ericbutera/amalgam/api/internal/metrics"
	"github.com/gin-gonic/gin"
)

func (s *server) routes() {
	s.router.GET("/", func(c *gin.Context) { c.Redirect(http.StatusMovedPermanently, internal.SwaggerUri) })
	s.router.GET("/health", func(c *gin.Context) {
		metrics.TestCounter.Inc()
		c.Status(http.StatusOK)
	})

	// TODO auth
	// auth := s.router.Group("/")
	// auth.Use(app.AuthRequired()) {
	s.router.GET("/swagger/*any", internal.Swagger())
	//}
}
