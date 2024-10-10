package server

import (
	"net/http"

	"github.com/ericbutera/amalgam/api/internal"
	"github.com/ericbutera/amalgam/api/internal/models"
	"github.com/ericbutera/amalgam/api/internal/service"
	"github.com/gin-gonic/gin"

	_ "github.com/ericbutera/amalgam/api/docs"
)

func (s *server) routes() {
	handlers := newHandlers(service.New(s.db) /*s.db*/)

	s.router.GET("/", func(c *gin.Context) { c.Redirect(http.StatusMovedPermanently, internal.SwaggerUri) })
	s.router.GET("/health", health)
	s.router.GET("/feeds", handlers.feeds)
	s.router.GET("/feed/:id/articles", handlers.articles)

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
	//db *gorm.DB
	svc *service.Service
}

func newHandlers(svc *service.Service /*db *gorm.DB*/) *handlers {
	return &handlers{
		//db: db
		svc: svc,
	}
}

// list feeds
// @Summary list feeds
// @Schemes
// @Description list feeds
// @Accept json
// @Produce json
// @Success 200 {object} FeedsResponse
// @Failure 500 {object} map[string]string
// @Router /feeds [get]
func (h *handlers) feeds(c *gin.Context) {
	feeds, err := h.svc.Feeds()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "unable to fetch feeds"})
	}

	c.JSON(http.StatusOK, FeedsResponse{
		Feeds: feeds,
	})
}

type FeedsResponse struct {
	Feeds []models.Feed `json:"feeds"`
}

// list articles for a feed
// @Summary list articles for a feed
// @Schemes
// @Description list articles for a feed
// @Accept json
// @Produce json
// @Param id path int true "Feed ID"
// @Success 200 {object} FeedArticlesResponse
// @Failure 500 {object} map[string]string
// @Router /feed/{id}/articles [get]
func (h *handlers) articles(c *gin.Context) {
	// TODO: simplify flow control
	// TODO: handle bad request
	// TODO: handle non-existent feed
	id := c.Param("id")

	articles, err := h.svc.GetArticlesByFeed(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "unable to fetch articles"})
	}

	c.JSON(http.StatusOK, FeedArticlesResponse{
		Articles: articles,
	})
}

type FeedArticlesResponse struct {
	Articles []models.Article `json:"articles"`
}
