package graph

import (
	"github.com/ericbutera/amalgam/graph/internal/config"
	"github.com/ericbutera/amalgam/pkg/client"
	pb "github.com/ericbutera/amalgam/pkg/feeds/v1"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	config    *config.Config
	apiClient *client.APIClient
	rpcClient pb.FeedServiceClient
}

func NewResolver(config *config.Config, apiClient *client.APIClient, rpcClient pb.FeedServiceClient) *Resolver {
	return &Resolver{
		config:    config,
		apiClient: apiClient,
		rpcClient: rpcClient,
	}
}

func NewApiClient(scheme string, host string) *client.APIClient {
	cfg := client.NewConfiguration()
	cfg.Scheme = scheme
	cfg.Host = host
	cfg.Debug = true
	return client.NewAPIClient(cfg)
}
