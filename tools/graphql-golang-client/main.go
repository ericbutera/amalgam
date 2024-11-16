package main

// TODO: convert to CLI

import (
	"log/slog"

	"github.com/Khan/genqlient/generate"
)

func main() {
	slog.Info("generating golang graphql client")
	generate.Main()
}
