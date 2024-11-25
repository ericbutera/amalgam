package server_test

import (
	"bytes"
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ericbutera/amalgam/services/graph/internal/config"
	"github.com/ericbutera/amalgam/services/graph/internal/server"
	"github.com/stretchr/testify/require"
)

func Test_ComplexityLimit_Exceeded(t *testing.T) {
	ctx := context.TODO()
	config := &config.Config{
		ComplexityLimit: 1,
		Port:            "9999",
	}
	srv, err := server.New(ctx, config, nil, nil)
	require.NoError(t, err)

	complexQuery := `{
		"operationName": "Feeds",
		"query": "query Feeds { feeds { id url name } feed(id: \"1\") { id url name } articles(feedId: \"1\") { id feedId url title imageUrl content description preview guid authorName authorEmail } }"
	}`
	req := httptest.NewRequest(http.MethodPost, "/query", bytes.NewBufferString(complexQuery))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	srv.Handler.ServeHTTP(rec, req)

	require.Equal(t, http.StatusOK, rec.Code)

	body := rec.Body.String()
	require.Contains(t, body, "COMPLEXITY_LIMIT_EXCEEDED")
}
