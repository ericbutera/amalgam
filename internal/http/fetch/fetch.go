package fetch

import (
	"context"
	"io"
	"log/slog"
	"net/http"
	"time"

	"github.com/ericbutera/amalgam/internal/http/transport"
)

const FetchTimeout = 10 * time.Second // TODO: WithTimeout opt

type CallbackParams struct {
	Reader      io.Reader
	Size        int64
	ContentType string
}

type Callback = func(params CallbackParams) error

func Url(ctx context.Context, url string, fetchCb Callback) error {
	// note: do not retry; workflow will handle retries
	// TODO: otel
	client := &http.Client{
		Transport: transport.NewLoggingTransport(
			transport.WithLogger(slog.Default()),
		),
		Timeout: FetchTimeout,
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return err
	}
	req.Header.Set("User-Agent", "curl/7.79.1")
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return fetchCb(CallbackParams{
		ContentType: resp.Header.Get("Content-Type"),
		Reader:      resp.Body,
		Size:        resp.ContentLength,
	})
}
