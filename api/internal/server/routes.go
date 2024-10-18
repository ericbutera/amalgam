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

	// TODO: require auth
	// TODO: move to static file server
	s.router.GET("/swagger/*any", internal.Swagger())

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
// @Param id path int true "Feed ID" minimum(1)
// @Success 200 {object} FeedResponse
// @Failure 500 {object} ErrorResponse
// @Router /feeds/{id} [get]
func (h *handlers) feedGet(c *gin.Context) {
	feed, err := h.svc.GetFeed(c.Request.Context(), c.Param("id"))
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
	var req CreateFeedRequest
	if err := c.Bind(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	feed := models.Feed{
		Url: req.Feed.Url,
	}
	if err := h.svc.CreateFeed(c.Request.Context(), &feed); err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: "unable to create feed"})
		return
	}
	c.JSON(http.StatusOK, FeedCreateResponse{
		Feed: &feed,
	})
}

// TODO: separate api from db
type CreateFeed struct {
	Url string `json:"url" binding:"required" example:"https://example.com/feed.xml"`
}
type CreateFeedRequest struct {
	Feed CreateFeed `json:"feed"`
}
type FeedCreateResponse struct {
	Feed *models.Feed `json:"feed"` // TODO: limit fields
}

// update feed
// @Summary update feed
// @Schemes
// @Description update feed
// @Accept json
// @Produce json
// @Param id path int true "Feed ID"
// @Success 200 {object} FeedUpdateResponse
// @Failure 500 {object} map[string]string
// @Router /feeds/{id} [post]
func (h *handlers) feedUpdate(c *gin.Context) {
	var req UpdateFeedRequest
	if err := c.Bind(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	feed := models.Feed{
		Url: req.Feed.Url,
	}
	if err := h.svc.UpdateFeed(c.Request.Context(), c.Param("id"), &feed); err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: "unable to create feed"})
		return
	}
	c.JSON(http.StatusOK, FeedUpdateResponse{
		Feed: &feed,
	})
}

// TODO: separate api from db
type UpdateFeed struct {
	ID  uint   `json:"id" binding:"required" example:"1"`
	Url string `json:"url" binding:"required" example:"https://example.com/feed.xml"`
}

type UpdateFeedRequest struct {
	Feed UpdateFeed `json:"feed"`
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
	feeds, err := h.svc.Feeds(c.Request.Context())
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
	article, err := h.svc.GetArticle(c.Request.Context(), c.Param("id"))
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
	articles, err := h.svc.GetArticlesByFeed(c.Request.Context(), c.Param("id"))
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
