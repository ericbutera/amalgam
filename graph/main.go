package main

//"github.com/99designs/gqlgen-contrib/prometheus"
import (
	"context"
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

	cfg := lo.Must(cfg.NewFromEnv[config.Config]())

	shutdown := lo.Must(otel.Setup(ctx))
	defer lo.Must0(shutdown(ctx))

	gql_prom.Register()

	c := lo.Must(rpc.NewClient(cfg.RpcHost, cfg.RpcInsecure))
	defer c.Conn.Close()

	srv := newServer(cfg, c.Client)
	srv.Use(otelgqlgen.Middleware())
	srv.Use(gql_prom.Tracer{})
	// TODO: complexity limit srv.Use(extension.ComplexityLimit{})

	router := newRouter(cfg.CorsAllowOrigins)
	router.Handle("/metrics", promhttp.Handler())
	router.Handle("/query", srv)
	router.Handle("/", playground.Handler("GraphQL playground", "/query"))

	slog.Info("running graphql playground", "port", cfg.Port)
	if err := http.ListenAndServe(":"+cfg.Port, router); err != nil {
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
