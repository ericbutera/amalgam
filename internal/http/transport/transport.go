package transport

import (
	"log/slog"
	"net/http"
)

// https://www.piotrbelina.com/blog/http-log/
type LoggingTransport struct {
	rt             http.RoundTripper
	logger         *slog.Logger
	detailedTiming bool
}

func (t *LoggingTransport) RoundTrip(r *http.Request) (*http.Response, error) {
	// do before request is sent, ex. start timer, log request
	resp, err := t.rt.RoundTrip(r)
	// do after the response is received, ex. end timer, log response
	return resp, err
}

type Option func(transport *LoggingTransport)

func NewLoggingTransport(options ...Option) *LoggingTransport {
	t := &LoggingTransport{
		rt:             http.DefaultTransport,
		logger:         slog.Default(),
		detailedTiming: false,
	}

	for _, option := range options {
		option(t)
	}

	return t
}

func WithRoundTripper(rt http.RoundTripper) Option {
	return func(t *LoggingTransport) {
		t.rt = rt
	}
}

func WithLogger(logger *slog.Logger) Option {
	return func(t *LoggingTransport) {
		t.logger = logger
	}
}
