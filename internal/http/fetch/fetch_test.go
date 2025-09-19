package fetch_test

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"testing"
	"time"

	"github.com/ericbutera/amalgam/internal/http/fetch"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type mockTransport struct {
	resp *http.Response
	err  error
}

func (m *mockTransport) RoundTrip(_ *http.Request) (*http.Response, error) {
	return m.resp, m.err
}

func TestUrl_Success(t *testing.T) {
	t.Parallel()

	resp := &http.Response{
		StatusCode:    http.StatusOK,
		Body:          io.NopCloser(bytes.NewBufferString("test data")),
		ContentLength: 9,
		Header:        http.Header{"Content-Type": []string{"text/plain"}},
	}

	client := &http.Client{
		Transport: &mockTransport{
			resp: resp,
		},
		Timeout: fetch.FetchTimeout,
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	fetcher, err := fetch.New(fetch.WithClient(client))
	require.NoError(t, err)

	err = fetcher.Url(ctx, "http://example.com", func(params fetch.CallbackParams) error {
		assert.Equal(t, resp.Header.Get("Content-Type"), params.ContentType)
		assert.Equal(t, resp.ContentLength, params.Size)
		body, err := io.ReadAll(params.Reader)
		require.NoError(t, err)
		assert.Equal(t, "test data", string(body))

		return nil
	}, nil)

	require.NoError(t, err)
}

// TODO: test etag
