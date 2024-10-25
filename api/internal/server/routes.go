package server

import (
	"net/http"

	"github.com/Khan/genqlient/graphql"
	"github.com/ericbutera/amalgam/api/internal"
	"github.com/ericbutera/amalgam/internal/db/models"
	"github.com/ericbutera/amalgam/internal/service"
	"github.com/gin-gonic/gin"

	graph_client "github.com/ericbutera/amalgam/graph/pkg/client"

	_ "github.com/ericbutera/amalgam/api/docs"
)

// TODO: do not show raw errors to the user

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
	handlers := newHandlers(s.service, s.graphClient)

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
	svc         *service.Service
	graphClient graphql.Client
}

func newHandlers(svc *service.Service, graphClient graphql.Client) *handlers {
	return &handlers{
		svc:         svc,
		graphClient: graphClient,
	}
}

type ErrorResponse struct {
	Error string `json:"error" example:"unable to fetch feeds"`
}

func (h *handlers) health(c *gin.Context) {
	c.Status(http.StatusOK)
}

// @Summary view feed
// @Schemes
// @Description view feed
// @Accept json
// @Produce json
// @Param id path string true "Feed ID"
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

// @Summary create feed
// @Schemes
// @Description create feed
// @Accept json
// @Produce json
// @Param request body CreateFeedRequest true "feed data"
// @Success 200 {object} FeedCreateResponse
// @Failure 500 {object} map[string]string
// @Router /feeds [post]
func (h *handlers) feedCreate(c *gin.Context) {
	var req CreateFeedRequest
	if err := c.Bind(&req); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: err.Error()})
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
		Id: feed.ID,
	})
}

// TODO: separate api from db
type CreateFeed struct {
	Name string `json:"name" binding:"omitempty" example:"My Feed"`
	Url  string `json:"url" binding:"required,url" example:"https://example.com/feed.xml"`
}
type CreateFeedRequest struct {
	Feed CreateFeed `json:"feed"`
}
type FeedCreateResponse struct {
	Id string `json:"id"`
}

// @Summary update feed
// @Schemes
// @Description update feed
// @Accept json
// @Produce json
// @Param id path string true "Feed ID"
// @Param request body UpdateFeedRequest true "feed data"
// @Success 200 {object} FeedUpdateResponse
// @Failure 500 {object} map[string]string
// @Router /feeds/{id} [put]
func (h *handlers) feedUpdate(c *gin.Context) {
	var req UpdateFeedRequest
	if err := c.Bind(&req); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: err.Error()})
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
	Name string `json:"name" binding:"omitempty" example:"My Feed"`
	Url  string `json:"url" binding:"required,url" example:"https://example.com/feed.xml"`
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

// @Summary list feeds
// @Schemes
// @Description list feeds
// @Accept json
// @Produce json
// @Success 200 {object} FeedsResponse
// @Failure 500 {object} map[string]string
// @Router /feeds [get]
func (h *handlers) feedsList(c *gin.Context) {
	resp, err := graph_client.Feeds(c.Request.Context(), h.graphClient)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{
			Error: err.Error(),
		})
	}

	feeds := []ListFeed{}
	for _, feed := range resp.Feeds {
		feeds = append(feeds, ListFeed{
			Id:   feed.Id,
			Name: feed.Name,
			Url:  feed.Url,
		})
	}
	c.JSON(http.StatusOK, FeedsResponse{
		Feeds: feeds,
	})
}

type ListFeed struct {
	Id   string `json:"id" example:"aaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaaaaaa"`
	Name string `json:"name" example:"Example"`
	Url  string `json:"url" example:"https://example.com/"`
}

type FeedsResponse struct {
	Feeds []ListFeed `json:"feeds"`
}

// @Summary view article
// @Schemes
// @Description view article
// @Accept json
// @Produce json
// @Param id path string true "Article ID"
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

// @Summary list articles for a feed
// @Schemes
// @Description list articles for a feed
// @Accept json
// @Produce json
// @Param id path string true "Feed ID"
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
