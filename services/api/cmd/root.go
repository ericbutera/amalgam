package cmd

import (
	"log/slog"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "api",
	Short: "api",
	Long:  "api",
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		slog.ErrorContext(rootCmd.Context(), "error running root command")
	}
}

func init() { //nolint:gochecknoinits
	rootCmd.AddCommand(NewServerCmd())
}
