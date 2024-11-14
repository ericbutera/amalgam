package grpc_test

import (
	"testing"

	"github.com/ericbutera/amalgam/rpc/internal/server/grpc"
	"github.com/stretchr/testify/assert"
)

func TestNewServer(t *testing.T) {
	srv := grpc.NewServer(nil, nil)
	assert.NotNil(t, srv)
}
