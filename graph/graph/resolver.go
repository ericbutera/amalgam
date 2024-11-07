package graph

import (
	"github.com/ericbutera/amalgam/graph/internal/config"
	pb "github.com/ericbutera/amalgam/pkg/feeds/v1"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	config    *config.Config
	rpcClient pb.FeedServiceClient
}

func NewResolver(config *config.Config, rpcClient pb.FeedServiceClient) *Resolver {
	return &Resolver{
		config:    config,
		rpcClient: rpcClient,
	}
}
