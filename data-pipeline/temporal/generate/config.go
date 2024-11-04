package generate

import "github.com/spf13/viper"

type Config struct {
	WorkflowID   string `mapstructure:"generate_workflow_id"`
	TaskQueue    string `mapstructure:"generate_task_queue"`
	GraphHost    string `mapstructure:"graph_host"`
	FakeHost     string `mapstructure:"fake_host"`
	RpcHost      string `mapstructure:"rpc_host" `
	RpcInsecure  bool   `mapstructure:"rpc_insecure"`
	TemporalHost string `mapstructure:"temporal_host"`
}

func init() {
	viper.SetDefault("generate_workflow_id", "generate-feeds")
	viper.SetDefault("generate_task_queue", "generate-feeds-queue")
	viper.SetDefault("fake_host", "")
	viper.SetDefault("graph_host", "")
	viper.SetDefault("rpc_host", "")
	viper.SetDefault("rpc_insecure", false)
	viper.SetDefault("temporal_host", "")
}
