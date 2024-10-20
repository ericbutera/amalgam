package main

import (
	"log/slog"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/ericbutera/amalgam/graph/graph"
	"github.com/ericbutera/amalgam/graph/internal/config"
)

func main() {
	config, err := config.NewConfigFromEnv()
	if err != nil {
		slog.Error("failed to load config", "error", err)
		os.Exit(1)
		return
	}

	apiClient := graph.NewApiClient(config.ApiScheme, config.ApiHost)
	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{
		Resolvers: graph.NewResolver(config, apiClient),
	}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	slog.Info("running graphql playground", "port", config.Port)
	if err := http.ListenAndServe(":"+config.Port, nil); err != nil {
		slog.Error("failed to start server", "error", err)
		os.Exit(1)
		return
	}
}
