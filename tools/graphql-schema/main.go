package main

// TODO: convert to CLI
// TODO: move to tools

import (
	"context"
	"fmt"
	"log/slog"
	"os"

	"github.com/suessflorian/gqlfetch"
)

var (
	serviceUrl = "http://localhost:8082/query"
	outfile    = "schema.graphql"
)

// Downloads the GraphQL schema from the locally running server.
// more info: https://github.com/Khan/genqlient/blob/main/docs/schema.md#fetching-your-schema
func main() {
	slog.Info("starting")

	schema, err := gqlfetch.BuildClientSchema(context.Background(), serviceUrl, false)
	if err != nil {
		slog.Error("unable to query graph service", "error", err)
		os.Exit(1)
	}

	if err = os.WriteFile(outfile, []byte(schema), 0644); err != nil {
		fmt.Println(err)
		slog.Error("unable to write schema", "error", err)
		os.Exit(1)
	}

	slog.Info("schema written", "outfile", outfile)
}
