package config

import "github.com/spf13/viper"

type Config struct {
	UseSchedule          bool   `mapstructure:"feed_use_schedule"`
	ScheduleID           string `mapstructure:"feed_schedule_id"`
	WorkflowID           string `mapstructure:"feed_workflow_id"`
	TaskQueue            string `mapstructure:"feed_task_queue"`
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
	viper.SetDefault("feed_use_schedule", false)
	viper.SetDefault("feed_schedule_id", "feed-schedule-id")
	viper.SetDefault("feed_workflow_id", "feed-workflow-id")
	viper.SetDefault("feed_task_queue", "feed-task-queue")
	viper.SetDefault("temporal_host", "")
	viper.SetDefault("rpc_host", "")
	viper.SetDefault("rpc_insecure", false)
	viper.SetDefault("minio_endpoint", "")
	viper.SetDefault("minio_access_key", "")
	viper.SetDefault("minio_secret_access_key", "")
	viper.SetDefault("minio_use_ssl", true)
	viper.SetDefault("minio_trace", false)
	viper.SetDefault("minio_region", "us-east-1")
}
