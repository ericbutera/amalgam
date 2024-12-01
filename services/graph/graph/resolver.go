package graph

import (
	"github.com/ericbutera/amalgam/internal/converters"
	"github.com/ericbutera/amalgam/internal/tasks"
	pb "github.com/ericbutera/amalgam/pkg/feeds/v1"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.
const (
	DefaultLimit int32 = 25
	LimitMax     int32 = 100
)

type Resolver struct {
	rpcClient pb.FeedServiceClient
	tasks     tasks.Tasks
	converter converters.Converter
	// TODO: auth authprovider
}

func NewResolver(rpcClient pb.FeedServiceClient, tasks tasks.Tasks) *Resolver {
	return &Resolver{
		rpcClient: rpcClient,
		tasks:     tasks,
		converter: converters.New(),
		// TODO: auth authprovider
	}
}
