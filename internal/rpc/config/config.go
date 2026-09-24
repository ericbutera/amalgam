package config

type Config struct {
	OtelEnable       bool   `env:"OTEL_ENABLE"        envDefault:"false"`                       // enable opentelemetry
	IgnoredSpanNames string `env:"IGNORED_SPAN_NAMES" envDefault:"grpc.health.v1.Health/Check"` // span names to ignore
	Port             string `env:"PORT"               envDefault:"8080"`                        // grpc server port
	MetricAddress    string `env:"METRIC_ADDRESS"     envDefault:":9090"`                       // metric server address
}
