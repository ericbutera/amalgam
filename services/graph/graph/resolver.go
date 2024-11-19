package graph

import (
	pb "github.com/ericbutera/amalgam/pkg/feeds/v1"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	rpcClient pb.FeedServiceClient
}

func NewResolver(rpcClient pb.FeedServiceClient) *Resolver {
	return &Resolver{
		rpcClient: rpcClient,
	}
}
