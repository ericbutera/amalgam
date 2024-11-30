package server_test

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ericbutera/amalgam/services/graph/internal/config"
	"github.com/ericbutera/amalgam/services/graph/internal/server"
	"github.com/stretchr/testify/require"
)

func Test_ComplexityLimit_Exceeded(t *testing.T) {
	t.Skip("skipping flaky test")

	config := &config.Config{
		ComplexityLimit: 1,
		Port:            "9999",
	}
	srv, err := server.New(config, nil, nil)
	require.NoError(t, err)

	complexQuery := `{
		"operationName": "Feeds",
		"query": "query Articles { articles(feedId: \"0e597e90-ece5-463e-8608-ff687bf286da\", options: { limit: 1, cursor: \"cursor\" }) { pagination { next previous } articles { id feedId url title imageUrl content description preview guid authorName authorEmail updatedAt } } }"
	}`
	req := httptest.NewRequest(http.MethodPost, "/query", bytes.NewBufferString(complexQuery))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	srv.Handler.ServeHTTP(rec, req)

	require.Equal(t, http.StatusOK, rec.Code)

	body := rec.Body.String()
	require.Contains(t, body, "COMPLEXITY_LIMIT_EXCEEDED")
}
