package main

// TODO: convert to CLI
// TODO: move to tools

import (
	"log/slog"

	"github.com/Khan/genqlient/generate"
)

func main() {
	slog.Info("generating")
	generate.Main()
}
