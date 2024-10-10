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
// @Success 200 {object} FeedsResponse
// @Failure 500 {object} map[string]string
// @Router /feeds [get]
func (h *handlers) feeds(c *gin.Context) {
	var feeds []models.Feed
	result := h.db.Find(&feeds).Limit(100)
	if result.Error != nil {
		// TODO error logging middleware
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
	var feed models.Feed
	// TODO: simplify flow control
	// TODO: handle bad request
	id := c.Param("id")
	result := h.db.First(&feed, id)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "feed not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "unable to fetch feeds"})
	}

	// TODO: create specific article listing which excludes feed obj & conent
	var articles []models.Article
	res := h.db.Find(&articles, "feed_id = ?", feed.ID).Limit(100)
	if res.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "unable to fetch articles"})
	}

	c.JSON(http.StatusOK, FeedArticlesResponse{
		Articles: articles,
	})
}

type FeedArticlesResponse struct {
	Articles []models.Article `json:"articles"`
}
