package app

import (
	"io"
	"log/slog"
	"net/http"
	"os"

	"go.temporal.io/sdk/client"
)

func NewTemporalClient(host string) (client.Client, error) {
	h := slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug})
	slog.SetDefault(slog.New(h))
	slog.SetLogLoggerLevel(slog.LevelDebug) // TODO: move to config

	opts := client.Options{
		Logger:   slog.Default(),
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
	client := &http.Client{}
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
