package feed_tasks

import "github.com/spf13/viper"

type Config struct {
	GenerateCount int    `mapstructure:"feed_tasks_count"`
	WorkflowID    string `mapstructure:"feed_tasks_workflow_id"`
	TaskQueue     string `mapstructure:"feed_tasks_task_queue"`
	GraphHost     string `mapstructure:"graph_host"`
	FakeHost      string `mapstructure:"fake_host"`
	RpcHost       string `mapstructure:"rpc_host"`
	RpcInsecure   bool   `mapstructure:"rpc_insecure"`
	TemporalHost  string `mapstructure:"temporal_host"`
}

func init() { //nolint:gochecknoinits
	viper.SetDefault("feed_tasks_count", 25)
	viper.SetDefault("feed_tasks_workflow_id", "feed-tasks-feeds")
	viper.SetDefault("feed_tasks_task_queue", "feed-tasks-feeds-queue")
	viper.SetDefault("fake_host", "")
	viper.SetDefault("graph_host", "")
	viper.SetDefault("rpc_host", "")
	viper.SetDefault("rpc_insecure", false)
	viper.SetDefault("temporal_host", "")
}
