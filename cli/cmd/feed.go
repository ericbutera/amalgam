package cmd

import (
	"fmt"
	"log/slog"
	"net/http"
	"net/url"

	"github.com/Khan/genqlient/graphql"
	graph_client "github.com/ericbutera/amalgam/graph/pkg/client"
	"github.com/spf13/cobra"
)

func NewFeedCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "feed",
		Short: "feed",
		Long:  "feed",
	}
}

func newGraphClient() graphql.Client {
	httpClient := http.Client{}
	return graphql.NewClient(
		"http://localhost:8082/query", // TODO: make this configurable
		&httpClient,
	)
}

func NewFeedListCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "list",
		Short: "list feeds",
		Run: func(cmd *cobra.Command, args []string) {
			res, err := graph_client.Feeds(cmd.Context(), newGraphClient())
			if err != nil {
				slog.Error("failed to list feeds", "error", err)
				return
			}
			slog.Debug("feed count", "feeds", len(res.Feeds))
			for _, feed := range res.Feeds {
				fmt.Printf("feed id: %s, url: %s\n", feed.Id, feed.Url)
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
