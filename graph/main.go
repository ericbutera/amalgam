package main

import (
	"log/slog"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/ericbutera/amalgam/graph/graph"
	"github.com/ericbutera/amalgam/graph/internal/config"
	cfg "github.com/ericbutera/amalgam/pkg/config"
)

func main() {
	cfg, err := cfg.NewFromEnv[config.Config]()
	if err != nil {
		slog.Error("failed to load config", "error", err)
		os.Exit(1)
		return
	}

	apiClient := graph.NewApiClient(cfg.ApiScheme, cfg.ApiHost)
	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{
		Resolvers: graph.NewResolver(cfg, apiClient),
	}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	slog.Info("running graphql playground", "port", cfg.Port)
	if err := http.ListenAndServe(":"+cfg.Port, nil); err != nil {
		slog.Error("failed to start server", "error", err)
		os.Exit(1)
		return
	}
}
