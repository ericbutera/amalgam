package config

import "github.com/spf13/viper"

type Config struct {
	UseSchedule          bool   `mapstructure:"use_schedule"`
	ScheduleID           string `mapstructure:"schedule_id"`
	WorkflowID           string `mapstructure:"workflow_id"`
	TaskQueue            string `mapstructure:"task_queue"`
	TemporalHost         string `mapstructure:"temporal_host"`
	RpcHost              string `mapstructure:"rpc_host"`
	RpcInsecure          bool   `mapstructure:"rpc_insecure"`
	MinioEndpoint        string `mapstructure:"minio_endpoint"`
	MinioAccessKey       string `mapstructure:"minio_access_key"`
	MinioSecretAccessKey string `mapstructure:"minio_secret_access_key"`
	MinioUseSsl          bool   `mapstructure:"minio_use_ssl"`
	MinioTrace           bool   `mapstructure:"minio_trace"`
	MinioRegion          string `mapstructure:"minio_region"`
}

func init() {
	viper.SetDefault("use_schedule", true)
	viper.SetDefault("temporal_host", "")
	viper.SetDefault("rpc_host", "")
	viper.SetDefault("rpc_insecure", false)
	viper.SetDefault("schedule_id", "feed-schedule-id")
	viper.SetDefault("workflow_id", "feed-workflow-id")
	viper.SetDefault("task_queue", "feed-task-queue")
	viper.SetDefault("minio_endpoint", "localhost:9001")
	viper.SetDefault("minio_access_key", "")
	viper.SetDefault("minio_secret_access_key", "")
	viper.SetDefault("minio_use_ssl", true)
	viper.SetDefault("minio_trace", false)
	viper.SetDefault("minio_region", "us-east-1")
}
