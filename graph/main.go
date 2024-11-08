package main

// github.com/99designs/gqlgen-contrib/prometheus
import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	gql_prom "github.com/ericbutera/amalgam/graph/extensions/prometheus"
	"github.com/ericbutera/amalgam/graph/graph"
	"github.com/ericbutera/amalgam/graph/internal/config"
	"github.com/ericbutera/amalgam/internal/http/server"
	"github.com/ericbutera/amalgam/internal/logger"
	"github.com/ericbutera/amalgam/pkg/config/env"
	pb "github.com/ericbutera/amalgam/pkg/feeds/v1"
	"github.com/ericbutera/amalgam/pkg/otel"
	rpc "github.com/ericbutera/amalgam/rpc/pkg/client"
	"github.com/go-chi/chi"
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

	shutdown := lo.Must(otel.Setup(ctx))
	defer lo.Must0(shutdown(ctx))

	gql_prom.Register()

	c := lo.Must(rpc.NewClient(config.RpcHost, config.RpcInsecure))
	defer c.Conn.Close()

	srv := newServer(config, c.Client)
	srv.Use(otelgqlgen.Middleware())
	srv.Use(gql_prom.Tracer{})
	// TODO: complexity limit srv.Use(extension.ComplexityLimit{})

	router := newRouter(config.CorsAllowOrigins)
	router.Handle("/metrics", promhttp.Handler())
	router.Handle("/query", srv)
	router.Handle("/", playground.Handler("GraphQL playground", "/query"))

	slog.Info("running graphql playground", "port", config.Port)
	server := lo.Must(server.New(
		server.WithAddr(":"+config.Port),
		server.WithHandler(router),
	))
	if err := server.ListenAndServe(); err != nil {
		slog.Error("failed to start server", "error", err)
		os.Exit(1)
		return
	}
}

func newServer(config *config.Config, rpcClient pb.FeedServiceClient) *handler.Server {
	return handler.NewDefaultServer(
		graph.NewExecutableSchema(graph.Config{
			Resolvers: graph.NewResolver(config, rpcClient),
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
		Debug:            true,
	}).Handler)

	return router
}
