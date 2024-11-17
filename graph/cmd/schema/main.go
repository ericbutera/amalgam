// minimal server that serves only the graphql schema
// intended to be used with generate schema
package main

import (
	"log/slog"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/ericbutera/amalgam/graph/graph"
	"github.com/ericbutera/amalgam/internal/http/server"
	"github.com/go-chi/chi"
	"github.com/samber/lo"
)

const (
	DefaultPort  = ":8082"
	DefaultRoute = "/query"
)

func main() {
	addr := os.Getenv("GRAPH_ADDR")
	if addr == "" {
		addr = ":8082"
	}
	srv := handler.NewDefaultServer(
		graph.NewExecutableSchema(graph.Config{
			Resolvers: graph.NewResolver(nil),
		}),
	)

	router := chi.NewRouter()
	router.Handle(DefaultRoute, srv)

	server := lo.Must(server.New(
		server.WithAddr(addr),
		server.WithHandler(router),
	))
	slog.Info("running server", "addr", addr, "route", DefaultRoute)
	if err := server.ListenAndServe(); err != nil {
		slog.Error("graph server error", "error", err)
		return
	}
}
