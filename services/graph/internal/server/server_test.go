package server_test

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ericbutera/amalgam/pkg/feeds/v1"
	"github.com/ericbutera/amalgam/services/graph/internal/config"
	"github.com/ericbutera/amalgam/services/graph/internal/server"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func Test_ComplexityLimit_Exceeded(t *testing.T) {
	config := &config.Config{
		ComplexityLimit: 1,
		Port:            "9999",
	}
	srv, err := server.New(config, nil, nil)
	require.NoError(t, err)

	query := `
		query Articles {
			articles(
				feedId: "0e597e90-ece5-463e-8608-ff687bf286da",
				options: { limit: 10, cursor: "0e597e90-ece5-463e-8608-ff687bf286da" }
			) {
				articles {
					id
					feedId
					url
					title
					imageUrl
					content
					description
					preview
					guid
					authorName
					authorEmail
					updatedAt
				}
			}
		}
	`
	complexQuery := fmt.Sprintf(`{ "operationName": "Articles", "query": %q }`, query)

	req := httptest.NewRequest(http.MethodPost, "/query", bytes.NewBufferString(complexQuery))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	srv.Handler.ServeHTTP(rec, req)

	require.Equal(t, http.StatusOK, rec.Code)

	body := rec.Body.String()
	require.Contains(t, body, "COMPLEXITY_LIMIT_EXCEEDED")
}

func TestReady(t *testing.T) {
	config := &config.Config{
		Port: "9999",
	}
	rpc := new(feeds.MockFeedServiceClient)
	srv, err := server.New(config, rpc, nil)
	require.NoError(t, err)

	rpc.EXPECT().Ready(mock.Anything, mock.Anything).Return(&feeds.ReadyResponse{}, nil)

	req := httptest.NewRequest(http.MethodGet, "/readyz", nil)
	rec := httptest.NewRecorder()
	srv.Handler.ServeHTTP(rec, req)

	require.Equal(t, http.StatusOK, rec.Code)
}

func TestHealth(t *testing.T) {
	config := &config.Config{
		Port: "9999",
	}
	srv, err := server.New(config, nil, nil)
	require.NoError(t, err)

	req := httptest.NewRequest(http.MethodGet, "/healthz", nil)
	rec := httptest.NewRecorder()
	srv.Handler.ServeHTTP(rec, req)

	require.Equal(t, http.StatusOK, rec.Code)
}

func TestMetrics(t *testing.T) {
	config := &config.Config{
		Port:       "9999",
		OtelEnable: true,
	}
	srv, err := server.New(config, nil, nil)
	require.NoError(t, err)

	req := httptest.NewRequest(http.MethodGet, "/metrics", nil)
	rec := httptest.NewRecorder()
	srv.Handler.ServeHTTP(rec, req)

	require.Equal(t, http.StatusOK, rec.Code)
}
