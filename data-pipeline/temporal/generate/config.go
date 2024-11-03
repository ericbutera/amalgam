package generate

import "github.com/spf13/viper"

type Config struct {
	GraphHost    string `mapstructure:"graph_host"`
	RpcHost      string `mapstructure:"rpc_host" `
	RpcInsecure  bool   `mapstructure:"rpc_insecure"`
	WorkflowID   string `mapstructure:"workflow_id"`
	TaskQueue    string `mapstructure:"task_queue"`
	TemporalHost string `mapstructure:"temporal_host"`
}

func init() {
	viper.SetDefault("graph_host", "")
	viper.SetDefault("rpc_host", "")
	viper.SetDefault("rpc_insecure", false)
	viper.SetDefault("workflow_id", "generate-feeds")
	viper.SetDefault("task_queue", "generate-feeds-queue")
	viper.SetDefault("temporal_host", "")
}
