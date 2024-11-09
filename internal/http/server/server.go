package server

import (
	"net/http"
	"time"
)

const (
	Address      = ":8080"
	ReadTimeout  = 10 * time.Second
	WriteTimeout = 10 * time.Second
	IdleTimeout  = 15 * time.Second
)

type Option func(*http.Server) error

func New(opts ...Option) (*http.Server, error) {
	server := &http.Server{
		Addr:         Address,
		Handler:      nil,
		ReadTimeout:  ReadTimeout,
		WriteTimeout: WriteTimeout,
		IdleTimeout:  IdleTimeout,
	}

	for _, opt := range opts {
		if err := opt(server); err != nil {
			return nil, err
		}
	}
	if server.Addr == "" {
		server.Addr = ":8080"
	}
	return server, nil
}

func WithAddr(addr string) Option {
	return func(s *http.Server) error {
		s.Addr = addr
		return nil
	}
}

func WithHandler(handler http.Handler) Option {
	return func(s *http.Server) error {
		s.Handler = handler
		return nil
	}
}

func WithReadTimeout(d time.Duration) Option {
	return func(s *http.Server) error {
		s.ReadTimeout = d
		return nil
	}
}

func WithWriteTimeout(d time.Duration) Option {
	return func(s *http.Server) error {
		s.WriteTimeout = d
		return nil
	}
}

func WithIdleTimeout(d time.Duration) Option {
	return func(s *http.Server) error {
		s.IdleTimeout = d
		return nil
	}
}
