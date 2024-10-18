package cmd

import (
	"fmt"
	"log/slog"
	"net/url"

	"github.com/ericbutera/amalgam/pkg/client"
	"github.com/spf13/cobra"
)

func NewFeedCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "feed",
		Short: "feed",
		Long:  "feed",
	}
}

func newClient() *client.APIClient {
	cfg := client.NewConfiguration()
	cfg.Scheme = "http"
	cfg.Host = "localhost:8080" // TODO: make this configurable
	return client.NewAPIClient(cfg)
}

func NewFeedListCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "list",
		Short: "list feeds",
		Run: func(cmd *cobra.Command, args []string) {
			res, http, err := newClient().
				DefaultAPI.
				FeedsGet(cmd.Context()).
				Execute()
			if err != nil {
				slog.Error("failed to list feeds", "error", err)
				return
			}
			slog.Debug("feed count", "feeds", len(res.Feeds), "code", http.StatusCode)
			for _, feed := range res.Feeds {
				fmt.Printf("feed id: %d, url: %s\n", feed.Id, feed.Url)
			}
		},
	}
}

func NewFeedAddCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "add [url]",
		Short: "add new feed",
		Args: func(cmd *cobra.Command, args []string) error {
			if err := cobra.ExactArgs(1)(cmd, args); err != nil {
				return err
			}
			u, err := url.Parse(args[0])
			if err == nil && u.Scheme != "" && u.Host != "" {
				return nil
			}
			return fmt.Errorf("invalid URL specified: %s", args[0])
		},
		Run: func(cmd *cobra.Command, args []string) {
			slog.Info("feed add!", "url", args[0])
			panic("TODO")
		},
	}
}
