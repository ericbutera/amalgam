package app

import (
	"io"
	"log/slog"
	"net/http"

	"github.com/ericbutera/amalgam/internal/http/transport"
	"github.com/ericbutera/amalgam/internal/logger"
	"go.temporal.io/sdk/client"
)

func NewTemporalClient(host string) (client.Client, error) {
	opts := client.Options{
		Logger:   logger.New(),
		HostPort: host,
	}
	c, err := client.Dial(opts)
	if err != nil {
		return nil, err
	}
	return c, nil
}

type FetchCallbackParams struct {
	Reader      io.Reader
	Size        int64
	ContentType string
}

type FetchCallback = func(params FetchCallbackParams) error

func FetchUrl(url string, fetchCb FetchCallback) error {
	// TODO: timeouts, retries, backoff, jitter
	client := &http.Client{
		Transport: transport.NewLoggingTransport(
			transport.WithLogger(slog.Default()),
		),
	}
	req, err := http.NewRequest("GET", url, nil)
	req.Header.Set("User-Agent", "curl/7.79.1")
	if err != nil {
		return err
	}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return fetchCb(FetchCallbackParams{
		ContentType: resp.Header.Get("Content-Type"),
		Reader:      resp.Body,
		Size:        resp.ContentLength,
	})
}
