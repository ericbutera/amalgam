package feed_tasks

import "github.com/spf13/viper"

const DefaultTasksCount = 25

type Config struct {
	GenerateCount int    `mapstructure:"feed_tasks_count"`
	WorkflowID    string `mapstructure:"feed_tasks_workflow_id"`
	TaskQueue     string `mapstructure:"feed_tasks_task_queue"` // queue name for feed_tasks.*Workflow
	FeedTaskQueue string `mapstructure:"feed_task_queue"`       // queue name for app.*Workflow
	GraphHost     string `mapstructure:"graph_host"`
	FakeHost      string `mapstructure:"fake_host"`
	RpcHost       string `mapstructure:"rpc_host"`
	RpcInsecure   bool   `mapstructure:"rpc_insecure"`
	TemporalHost  string `mapstructure:"temporal_host"`
}

func init() { //nolint:gochecknoinits
	viper.SetDefault("feed_tasks_count", DefaultTasksCount)
	viper.SetDefault("feed_tasks_workflow_id", "feed-tasks-feeds")
	viper.SetDefault("feed_tasks_task_queue", "feed-tasks-feeds-queue")
	viper.SetDefault("feed_task_queue", "feed-task-queue")
	viper.SetDefault("fake_host", "")
	viper.SetDefault("graph_host", "")
	viper.SetDefault("rpc_host", "")
	viper.SetDefault("rpc_insecure", false)
	viper.SetDefault("temporal_host", "")
}
