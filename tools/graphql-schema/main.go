package main

// TODO: convert to CLI

import (
	"context"
	"log/slog"
	"os"

	"github.com/suessflorian/gqlfetch"
)

const (
	serviceURL  = "http://localhost:8082/query" // TODO: config via env/cli param
	destination = "schema.graphql"
	fileMode    = 0o600
)

// Downloads the GraphQL schema from the locally running server.
// more info: https://github.com/Khan/genqlient/blob/main/docs/schema.md#fetching-your-schema
func main() {
	url := os.Getenv("GRAPH_HOST")
	if url == "" {
		url = serviceURL
	}

	slog.Info("generating golang graphql schema", "serviceURL", url)

	schema, err := gqlfetch.BuildClientSchema(context.Background(), url, false)
	if err != nil {
		slog.Error("unable to query graph service", "error", err)
		os.Exit(1)
	}

	if err = os.WriteFile(destination, []byte(schema), fileMode); err != nil {
		slog.Error("unable to write schema", "error", err)
		os.Exit(1)
	}

	slog.Info("schema written", "destination", destination)
}
