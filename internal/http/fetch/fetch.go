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

type ExtraParams struct {
	Etag string
}

// fetch content from a URL. intended to fetch RSS feeds.
type Fetch interface {
	Url(ctx context.Context, url string, fetchCb Callback, extra *ExtraParams) error
}

type CallbackParams struct {
	Reader      io.ReadCloser
	Size        int64
	ContentType string
	StatusCode  int
	Etag        string
}

type Callback = func(params CallbackParams) error

// TODO: ensure there is a cooldown on domain name
// TODO: handle etags (if fetch history contains etag, send it)
func Url(ctx context.Context, url string, fetchCb Callback, extra *ExtraParams) error {
	f, err := New()
	if err != nil {
		return err
	}

	return f.Url(ctx, url, fetchCb, extra)
}

type Http struct {
	client *http.Client
}

type Option func(*Http) error

func New(opts ...Option) (*Http, error) {
	f := &Http{}
	for _, opt := range opts {
		err := opt(f)
		if err != nil {
			return nil, err
		}
	}

	if f.client == nil {
		f.client = &http.Client{
			Transport: transport.NewLoggingTransport(
				transport.WithLogger(slog.Default()),
			),
			Timeout: FetchTimeout,
		}
	}

	return f, nil
}

func WithClient(c *http.Client) Option {
	return func(f *Http) error {
		f.client = c
		return nil
	}
}

func WithTransport(t http.RoundTripper) Option {
	return func(f *Http) error {
		f.client.Transport = t
		return nil
	}
}

func (f *Http) Url(ctx context.Context, url string, fetchCb Callback, extra *ExtraParams) error {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return err
	}

	req.Header.Set("User-Agent", "amalgam/1.4") // TODO: inject version

	if extra != nil && extra.Etag != "" {
		req.Header.Set("If-None-Match", extra.Etag)
	}

	resp, err := f.client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// TODO: this isn't the right spot, but each time fetch is called the system needs to save a fetch_history record.
	// - track etag to prevent refetching unchanged content (requires etag, fetch time)
	// - track last fetch to prevent refetching too soon
	// - track rate limits (429) and any "request after" time values - https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Retry-After
	// - content will also have different change cadence, should be tracked

	return fetchCb(CallbackParams{
		StatusCode:  resp.StatusCode,
		ContentType: resp.Header.Get("Content-Type"),
		Reader:      resp.Body,
		Size:        resp.ContentLength,
		Etag:        resp.Header.Get("Etag"),
	})
}
