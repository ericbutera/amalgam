package feed_tasks

type Config struct {
	GenerateCount  int    `env:"FEED_TASKS_COUNT"       envDefault:"25"`
	WorkflowID     string `env:"FEED_TASKS_WORKFLOW_ID" envDefault:"feed-tasks-feeds"`
	TaskQueue      string `env:"TASK_QUEUE"             envDefault:"feed-tasks-queue"` // queue name for feed_tasks.*Workflow
	FeedFetchQueue string `env:"FEED_FETCH_TASK_QUEUE"  envDefault:"feed-fetch-queue"` // feed fetch worker
	FeedAddQueue   string `env:"FEED_ADD_TASK_QUEUE"    envDefault:"feed-add-queue"`   // feed add worker
	GraphHost      string `env:"GRAPH_HOST"`
	FakeHost       string `env:"FAKE_HOST"`
	RpcHost        string `env:"RPC_HOST"`
	RpcInsecure    bool   `env:"RPC_INSECURE"           envDefault:"false"`
	TemporalHost   string `env:"TEMPORAL_HOST"`
}
