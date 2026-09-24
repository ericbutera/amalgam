package config

type Config struct {
	OtelEnable        bool   `env:"OTEL_ENABLE"            envDefault:"false"`
	IgnoredSpanNames  string `env:"IGNORED_SPAN_NAMES"`
	Port              string `env:"GRAPH_PORT"             envDefault:"8080"`
	ComplexityLimit   int    `env:"GRAPH_COMPLEXITY_LIMIT" envDefault:"20"`
	RpcHost           string `env:"RPC_HOST"               envDefault:"rpc:50051"`
	RpcInsecure       bool   `env:"RPC_INSECURE"           envDefault:"false"`
	CorsAllowOrigins  string `env:"CORS_ALLOW_ORIGINS"`
	CorsAllowMethods  string `env:"CORS_ALLOW_METHODS"`
	CorsAllowHeaders  string `env:"CORS_ALLOW_HEADERS"`
	CorsExposeHeaders string `env:"CORS_EXPOSE_HEADERS"`
}
