package server

import (
	"fmt"
	"log"
	"log/slog"
	"net"

	pb "github.com/ericbutera/amalgam/pkg/rpc/proto"
	"github.com/pkg/errors"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type Server struct {
	port     string
	srv      *grpc.Server
	listener net.Listener
	pb.UnimplementedFeedServiceServer
}

func (s *Server) Serve() {
	slog.Info("launching server", "port", s.port)
	if err := s.srv.Serve(s.listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

type Option func(*Server) error

func New(opts ...Option) (*Server, error) {
	server := Server{
		srv: grpc.NewServer(),
	}

	for _, opt := range opts {
		if err := opt(&server); err != nil {
			return nil, err
		}
	}

	if server.port == "" {
		return nil, fmt.Errorf("port is required")
	}
	if server.listener == nil {
		listener, err := net.Listen("tcp", fmt.Sprintf(":%s", server.port))
		if err != nil {
			return nil, errors.Errorf("failed to listen: %v", err)
		}
		server.listener = listener
	}

	pb.RegisterFeedServiceServer(server.srv, &Server{})
	reflection.Register(server.srv)

	return &server, nil
}

func WithPort(port string) Option {
	return func(s *Server) error {
		s.port = port
		return nil
	}
}
func WithListener(lis net.Listener) Option {
	return func(s *Server) error {
		s.listener = lis
		return nil
	}
}

// TODO: interceptors:
// - otel tracing
// - logging
// - metrics
