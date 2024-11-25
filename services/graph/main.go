package main

import (
	"context"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/ericbutera/amalgam/internal/http/server"
	"github.com/ericbutera/amalgam/internal/logger"
	"github.com/ericbutera/amalgam/internal/tasks"
	"github.com/ericbutera/amalgam/pkg/config/env"
	pb "github.com/ericbutera/amalgam/pkg/feeds/v1"
	"github.com/ericbutera/amalgam/pkg/otel"
	gql_prom "github.com/ericbutera/amalgam/services/graph/extensions/prometheus"
	"github.com/ericbutera/amalgam/services/graph/graph"
	"github.com/ericbutera/amalgam/services/graph/internal/config"
	rpc "github.com/ericbutera/amalgam/services/rpc/pkg/client"
	"github.com/go-chi/chi"
	"github.com/go-chi/httplog/v2"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/ravilushqa/otelgqlgen"
	"github.com/rs/cors"
	"github.com/samber/lo"
)

func main() {
	slog := logger.New()

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	config := lo.Must(env.New[config.Config]())

	client, closer := lo.Must2(rpc.New(config.RpcHost, config.RpcInsecure))
	defer func() { lo.Must0(closer()) }()

	tasks := lo.Must(tasks.NewTemporalFromEnv())
	defer tasks.Close()

	srv := newServer(client, tasks)

	router := newRouter(config.CorsAllowOrigins)
	router.Handle("/query", srv)
	router.Handle("/healthz", newHealthzHandler())
	router.Handle("/readyz", newReadyzHandler(client))
	router.Handle("/", playground.Handler("GraphQL playground", "/query"))

	if config.OtelEnable {
		shutdown := lo.Must(otel.Setup(ctx, config.IgnoredSpanNames))
		defer func() { lo.Must0(shutdown(ctx)) }()

		gql_prom.Register()
		srv.Use(otelgqlgen.Middleware())
		srv.Use(gql_prom.Tracer{})

		router.Handle("/metrics", promhttp.Handler())
	}

	srv.Use(extension.FixedComplexityLimit(config.ComplexityLimit))

	slog.Info("running graph", "port", config.Port)
	server := lo.Must(server.New(
		server.WithAddr(":"+config.Port),
		server.WithHandler(router),
	))

	// TODO: cancel context server.BaseContext

	go func() {
		<-ctx.Done()
		slog.Info("shutting down server")

		if err := server.Shutdown(ctx); err != nil {
			slog.Error("server shutdown failed", "error", err)
		}
	}()

	if err := server.ListenAndServe(); err != nil {
		slog.Error("failed to start server", "error", err)
		os.Exit(1) //nolint: gocritic
		return
	}
}

func newHealthzHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusOK)
	}
}

func newReadyzHandler(client pb.FeedServiceClient) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// TODO: add a low cost readiness endpoint (listfeeds could be expensive)
		_, err := client.ListFeeds(r.Context(), &pb.ListFeedsRequest{})
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
