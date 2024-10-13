package server

import (
	"net/http"

	"github.com/ericbutera/amalgam/api/internal"
	"github.com/ericbutera/amalgam/api/internal/models"
	"github.com/ericbutera/amalgam/api/internal/service"
	"github.com/gin-gonic/gin"

	_ "github.com/ericbutera/amalgam/api/docs"
)

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

func (s *server) routes() {
	handlers := newHandlers(s.service)

	s.router.GET("/", func(c *gin.Context) { c.Redirect(http.StatusMovedPermanently, internal.SwaggerUri) })
	s.router.GET("/health", handlers.health)
	s.router.GET("/swagger/*any", internal.Swagger()) // TODO: require auth

	v1 := s.router.Group("/v1")
	{
		// TODO auth: auth.Use(app.AuthRequired()) {
		v1.GET("/feeds", handlers.feedsList)
		v1.GET("/feeds/:id", handlers.feedGet)
		v1.POST("/feeds", handlers.feedCreate)
		v1.PUT("/feeds/:id", handlers.feedUpdate)
		v1.GET("/feeds/:id/articles", handlers.articles)
		v1.GET("/articles/:id", handlers.article)
	}

}

type handlers struct {
	svc *service.Service
}

func newHandlers(svc *service.Service /*db *gorm.DB*/) *handlers {
	return &handlers{
		svc: svc,
	}
}

type ErrorResponse struct {
	Error string `json:"error" example:"unable to fetch feeds"`
}

// Health check
// @Summary Health check
// @Schemes
// @Description Health check
// @Accept json
// @Produce json
// @Success 200
// @Router /health [get]
func (h *handlers) health(c *gin.Context) {
	c.Status(http.StatusOK)
}

// view feed
// @Summary view feed
// @Schemes
// @Description view feed
// @Accept json
// @Produce json
// @Example
// @Param id path int true "Feed ID" example(1) minimum(1)
// @Success 200 {object} FeedResponse
// @Failure 500 {object} ErrorResponse
// @Router /feeds/{id} [get]
func (h *handlers) feedGet(c *gin.Context) {
	feed, err := h.svc.GetFeed(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: "unable to get feed"})
	}

	c.JSON(http.StatusOK, FeedResponse{
		Feed: feed,
	})
}

type FeedResponse struct {
	Feed *models.Feed `json:"feed"`
}

// createFeed
// create feed
// @Summary create feed
// @Schemes
// @Description create feed
// @Accept json
// @Produce json
// @Success 200 {object} FeedCreateResponse
// @Failure 500 {object} map[string]string
// @Router /feeds [post]
func (h *handlers) feedCreate(c *gin.Context) {
	// normalize URL to prevent duplicates
	// create feed if not exists
	// create user_feed if not exists
	c.JSON(http.StatusNotImplemented, ErrorResponse{Error: "not implemented"})
}

type FeedCreateResponse struct {
	Feed *models.Feed `json:"feed"`
}

// update feed
// @Summary update feed
// @Schemes
// @Description update feed
// @Accept json
// @Produce json
// @Param id path int true "Feed ID" example(1)
// @Success 200 {object} FeedUpdateResponse
// @Failure 500 {object} map[string]string
// @Router /feeds/{id} [post]
func (h *handlers) feedUpdate(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, ErrorResponse{Error: "not implemented"})
}

type FeedUpdateResponse struct {
	Feed *models.Feed `json:"feed"`
}

// TODO feedDelete
// - delete from feeds_users
// - note: prevent fetch if feed not in feeds_users

// list feeds
// @Summary list feeds
// @Schemes
// @Description list feeds
// @Accept json
// @Produce json
// @Success 200 {object} FeedsResponse
// @Failure 500 {object} map[string]string
// @Router /feeds [get]
func (h *handlers) feedsList(c *gin.Context) {
	feeds, err := h.svc.Feeds()
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: "unable to fetch feeds"})
	}

	c.JSON(http.StatusOK, FeedsResponse{
		Feeds: feeds,
	})
}

type FeedsResponse struct {
	Feeds []models.Feed `json:"feeds"`
}

// view article
// @Summary view article
// @Schemes
// @Description view article
// @Accept json
// @Produce json
// @Param id path int true "Article ID"
// @Success 200 {object} ArticleResponse
// @Failure 500 {object} map[string]string
// @Router /articles/{id} [get]
func (h *handlers) article(c *gin.Context) {
	article, err := h.svc.GetArticle(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: "unable to fetch articles"})
	}

	c.JSON(http.StatusOK, ArticleResponse{
		Article: article,
	})
}

type ArticleResponse struct {
	Article *models.Article `json:"article"`
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
// @Router /feeds/{id}/articles [get]
func (h *handlers) articles(c *gin.Context) {
	articles, err := h.svc.GetArticlesByFeed(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: "unable to fetch articles"})
	}

	c.JSON(http.StatusOK, FeedArticlesResponse{
		Articles: articles,
	})
}

type FeedArticlesResponse struct {
	Articles []models.Article `json:"articles"`
}
