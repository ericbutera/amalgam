package config

import "github.com/spf13/viper"

type Config struct {
	TaskQueue            string `mapstructure:"task_queue"`
	TemporalHost         string `mapstructure:"temporal_host"`
	RpcHost              string `mapstructure:"rpc_host"`
	RpcInsecure          bool   `mapstructure:"rpc_insecure"`
	MinioEndpoint        string `mapstructure:"minio_endpoint"`
	MinioAccessKey       string `mapstructure:"minio_access_key"`
	MinioSecretAccessKey string `mapstructure:"minio_secret_access_key"`
	MinioUseSsl          bool   `mapstructure:"minio_use_ssl"`
}

func init() {
	viper.SetDefault("temporal_host", "")
	viper.SetDefault("rpc_host", "")
	viper.SetDefault("rpc_insecure", false)
	viper.SetDefault("task_queue", "feed-task-queue")
	viper.SetDefault("minio_endpoint", "localhost:9001")
	viper.SetDefault("minio_access_key", "")
	viper.SetDefault("minio_secret_access_key", "")
	viper.SetDefault("minio_use_ssl", true)
}
