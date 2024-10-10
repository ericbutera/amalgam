package server

import (
	"github.com/gin-gonic/gin"
)

const (
	MiddlewareName = "api"
)

type server struct {
	// Gin Router
	router *gin.Engine
}

func New() *server {
	// TODO: ensure gin uses signal.NotifyContext ctx
	// TODO: ensure gin uses slog as logger (for otel exporter)
	s := &server{router: gin.Default()}
	s.middleware()
	s.metrics()
	s.routes()
	return s
}

func (s *server) Run() error {
	return s.router.Run()
}

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
