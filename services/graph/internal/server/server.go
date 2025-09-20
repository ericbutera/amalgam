package server

import (
	"log/slog"
	"net/http"
	"strings"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/ericbutera/amalgam/internal/http/server"
	"github.com/ericbutera/amalgam/internal/tasks"
	pb "github.com/ericbutera/amalgam/pkg/feeds/v1"
	gql_prom "github.com/ericbutera/amalgam/services/graph/extensions/prometheus"
	"github.com/ericbutera/amalgam/services/graph/graph"
	"github.com/ericbutera/amalgam/services/graph/internal/config"
	"github.com/ericbutera/amalgam/services/graph/internal/middleware"
	"github.com/go-chi/chi"
	"github.com/go-chi/httplog/v2"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/ravilushqa/otelgqlgen"
	"github.com/rs/cors"
)

func New(config *config.Config, rpcClient pb.FeedServiceClient, tasks tasks.Tasks) (*http.Server, error) {
	srv := newServer(rpcClient, tasks)

	router := newRouter(strings.Split(config.CorsAllowOrigins, ","))
	router.Handle("/query", middleware.Auth(srv))
	router.Handle("/healthz", newHealthzHandler())
	router.Handle("/readyz", newReadyzHandler(rpcClient))
	router.Handle("/", playground.Handler("GraphQL playground", "/query"))

	if config.OtelEnable {
		gql_prom.Register()
		srv.Use(otelgqlgen.Middleware())
		srv.Use(gql_prom.Tracer{})

		router.Handle("/metrics", promhttp.Handler())
	}

	srv.Use(extension.FixedComplexityLimit(config.ComplexityLimit))
	srv.Use(middleware.NewErrorLogging())

	return server.New(
		server.WithAddr(":"+config.Port),
		server.WithHandler(router),
	)
}

func newHealthzHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusOK)
	}
}

func newReadyzHandler(client pb.FeedServiceClient) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_, err := client.Ready(r.Context(), &pb.ReadyRequest{})
		if err != nil {
			w.WriteHeader(http.StatusServiceUnavailable)
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}

func newServer(rpcClient pb.FeedServiceClient, tasks tasks.Tasks) *handler.Server {
	return handler.NewDefaultServer(
		graph.NewExecutableSchema(graph.Config{
			Resolvers: graph.NewResolver(rpcClient, tasks),
		}),
	)
}

func newRouter(allowOrigins []string) *chi.Mux {
	router := chi.NewRouter()
	// https://github.com/99designs/gqlgen/blob/master/docs/content/recipes/cors.md
	// Add CORS middleware around every request
	// read https://github.com/rs/cors for full option listing
	router.Use(cors.New(cors.Options{
		AllowedOrigins:   allowOrigins,
		AllowCredentials: true,
		Debug:            false,
	}).Handler)

	logger := httplog.NewLogger("graph", httplog.Options{
		LogLevel: slog.LevelDebug, // TODO: configurable
	})
	router.Use(httplog.RequestLogger(logger, []string{"/healthz", "/readyz", "/metrics"}))

	return router
}
