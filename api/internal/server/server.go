package server

import (
	"github.com/Khan/genqlient/graphql"
	"github.com/ericbutera/amalgam/api/internal/config"
	"github.com/gin-gonic/gin"
)

const (
	MiddlewareName = "api"
)

type server struct {
	config      *config.Config
	router      *gin.Engine
	graphClient graphql.Client
}

type ServerOption func(*server) error

func New(options ...ServerOption) (*server, error) {
	s := &server{
		router: gin.New(),
	}

	for _, o := range options {
		if err := o(s); err != nil {
			return nil, err
		}
	}

	s.middleware()
	s.metrics()
	s.routes()

	return s, nil
}

func WithConfig(cfg *config.Config) func(*server) error {
	return func(s *server) error {
		s.config = cfg
		return nil
	}
}

func WithGraphClient(client graphql.Client) func(*server) error {
	return func(s *server) error {
		s.graphClient = client
		return nil
	}
}

func (s *server) Run() error {
	return s.router.Run()
}
