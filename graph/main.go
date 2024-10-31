package main

//"github.com/99designs/gqlgen-contrib/prometheus"
import (
	"context"
	"errors"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	gql_prom "github.com/ericbutera/amalgam/graph/extensions/prometheus"
	"github.com/ericbutera/amalgam/graph/graph"
	"github.com/ericbutera/amalgam/graph/internal/config"
	"github.com/ericbutera/amalgam/internal/logger"
	cfg "github.com/ericbutera/amalgam/pkg/config"
	"github.com/ericbutera/amalgam/pkg/otel"
	rpc "github.com/ericbutera/amalgam/rpc/pkg/client"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/ravilushqa/otelgqlgen"
)

func main() {
	slog := logger.New()

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	cfg, err := cfg.NewFromEnv[config.Config]()
	if err != nil {
		slog.Error("failed to load config", "error", err)
		os.Exit(1)
		return
	}

	shutdown, err := otel.Setup(ctx)
	if err != nil {
		slog.Error("failed to setup otel", "error", err)
		os.Exit(1)
	}
	defer func() {
		err = errors.Join(err, shutdown(context.Background()))
	}()

	gql_prom.Register()

	c, err := rpc.NewClient(cfg.RpcHost, cfg.RpcInsecure)
	if err != nil {
		slog.Error("failed to connect to rpc server", "error", err)
		os.Exit(1)
	}
	defer c.Conn.Close()

	apiClient := graph.NewApiClient(cfg.ApiScheme, cfg.ApiHost)

	srv := handler.NewDefaultServer(
		graph.NewExecutableSchema(graph.Config{
			Resolvers: graph.NewResolver(cfg, apiClient, c.Client),
		}),
	)

	srv.Use(otelgqlgen.Middleware())
	srv.Use(gql_prom.Tracer{})
	// TODO: complexity limit srv.Use(extension.ComplexityLimit{})

	http.Handle("/metrics", promhttp.Handler())
	http.Handle("/query", srv)
	http.Handle("/", playground.Handler("GraphQL playground", "/query"))

	slog.Info("running graphql playground", "port", cfg.Port)
	if err := http.ListenAndServe(":"+cfg.Port, nil); err != nil {
		slog.Error("failed to start server", "error", err)
		os.Exit(1)
		return
	}
}
