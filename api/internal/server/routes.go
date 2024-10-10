package server

import (
	"net/http"

	"github.com/ericbutera/amalgam/api/internal"
	"github.com/gin-gonic/gin"

	_ "github.com/ericbutera/amalgam/api/docs"
)

func (s *server) routes() {
	s.router.GET("/", func(c *gin.Context) { c.Redirect(http.StatusMovedPermanently, internal.SwaggerUri) })
	s.router.GET("/health", health)

	// TODO auth
	// auth := s.router.Group("/")
	// auth.Use(app.AuthRequired()) {
	s.router.GET("/swagger/*any", internal.Swagger())
	//}
}

// Health check
// @Summary Health check
// @Schemes
// @Description Health check
// @Accept json
// @Produce json
// @Success 200
// @Router /health [get]
func health(c *gin.Context) {
	c.Status(http.StatusOK)
}
