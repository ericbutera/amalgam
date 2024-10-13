package server

import (
	"errors"

	"github.com/ericbutera/amalgam/api/internal/config"
	"github.com/ericbutera/amalgam/api/internal/db"
	"github.com/ericbutera/amalgam/api/internal/service"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

const (
	MiddlewareName = "api"
)

type server struct {
	config  *config.Config
	router  *gin.Engine
	db      *gorm.DB
	service *service.Service
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

	if s.db == nil {
		return nil, errors.New("database not set")
	}

	s.service = service.New(s.db)
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

func WithSqlite(name string) func(*server) error {
	return func(s *server) error {
		db, err := db.Sqlite(name)
		if err != nil {
			return err
		}
		s.db = db
		return nil
	}
}

func WithMysql(dsn string) func(*server) error {
	return func(s *server) error {
		db, err := db.Mysql(dsn)
		if err != nil {
			return err
		}
		s.db = db
		return nil
	}
}

func (s *server) Run() error {
	return s.router.Run()
}
