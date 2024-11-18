package grpc_test

import (
	"testing"

	"github.com/ericbutera/amalgam/services/rpc/internal/server/grpc"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewServer(t *testing.T) {
	srv, err := grpc.NewServer(nil, nil)
	require.NoError(t, err)
	assert.NotNil(t, srv)
}
