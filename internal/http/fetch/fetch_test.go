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

/*
func TestUrl(t *testing.T) {
	tests := []struct {
		name        string
		resp        *http.Response
		err         error
		callbackErr error
		expectedErr error
	}{
		{
			name: "successful fetch",
			resp: &http.Response{
				StatusCode:    http.StatusOK,
				Body:          io.NopCloser(bytes.NewBufferString("test data")),
				ContentLength: 9,
				Header:        http.Header{"Content-Type": []string{"text/plain"}},
			},
			expectedErr: nil,
		},
		{
			name:        "fetch error",
			err:         errors.New("fetch error"),
			expectedErr: errors.New("fetch error"),
		},
		// {
		// 	name: "callback error",
		// 	resp: &http.Response{
		// 		StatusCode:    http.StatusOK,
		// 		Body:          io.NopCloser(bytes.NewBufferString("test data")),
		// 		ContentLength: 9,
		// 		Header:        http.Header{"Content-Type": []string{"text/plain"}},
		// 	},
		// 	callbackErr: errors.New("callback error"),
		// 	expectedErr: errors.New("callback error"),
		// },
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client := &http.Client{
				Transport: &mockTransport{
					resp: tt.resp,
					err:  tt.err,
				},
				Timeout: fetch.FetchTimeout,
			}

			// Create a context with timeout to test cancellation.
			ctx, cancel := context.WithTimeout(context.Background(), time.Second)
			defer cancel()

			fetcher, err := fetch.New(fetch.WithClient(client))
			require.NoError(t, err)

			err = fetcher.Url(ctx, "http://example.com", func(params fetch.CallbackParams) error {
				if tt.callbackErr != nil {
					return tt.callbackErr
				}
				assert.Equal(t, tt.resp.Header.Get("Content-Type"), params.ContentType)
				assert.Equal(t, tt.resp.ContentLength, params.Size)
				body, err := io.ReadAll(params.Reader)
				require.NoError(t, err)
				assert.Equal(t, "test data", string(body))
				return nil
			})

			if tt.expectedErr != nil {
				require.Error(t, err)
				assert.Equal(t, tt.expectedErr.Error(), err.Error())
			} else {
				require.NoError(t, err)
			}
		})
	}
}
*/

func TestUrl_Success(t *testing.T) {
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
	})

	require.NoError(t, err)
}
