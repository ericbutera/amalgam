package config

type Config struct {
	OtelEnable        bool   `env:"OTEL_ENABLE"         envDefault:"false"`
	CorsAllowOrigins  string `env:"CORS_ALLOW_ORIGINS"`
	CorsAllowMethods  string `env:"CORS_ALLOW_METHODS"`
	CorsAllowHeaders  string `env:"CORS_ALLOW_HEADERS"`
	CorsExposeHeaders string `env:"CORS_EXPOSE_HEADERS"`
	GraphHost         string `env:"GRAPH_HOST"`
	// TODO:
	// log level
}
