package server

import (
	"net/http"

	"github.com/ericbutera/amalgam/api/internal"
	"github.com/ericbutera/amalgam/api/internal/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	_ "github.com/ericbutera/amalgam/api/docs"
)

func (s *server) routes() {
	handlers := newHandlers(s.db)

	s.router.GET("/", func(c *gin.Context) { c.Redirect(http.StatusMovedPermanently, internal.SwaggerUri) })
	s.router.GET("/health", health)
	s.router.GET("/feeds", handlers.feeds)

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

type handlers struct {
	db *gorm.DB
}

func newHandlers(db *gorm.DB) *handlers {
	return &handlers{db: db}
}

// list feeds
// @Summary list feeds
// @Schemes
// @Description list feeds
// @Accept json
// @Produce json
// @Success 200 {object} []models.Feed
// @Router /feeds [get]
func (h *handlers) feeds(c *gin.Context) {
	var feeds []models.Feed
	result := h.db.Find(&feeds).Limit(100)
	if result.Error != nil {
		// TODO error logging middleware
		c.JSON(http.StatusInternalServerError, gin.H{"error": "unable to fetch feeds"})
	}

	c.JSON(http.StatusOK, gin.H{
		"feeds": feeds,
	})
}
