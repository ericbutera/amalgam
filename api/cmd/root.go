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
	if err := rootCmd.Execute(); err != nil {
		slog.ErrorContext(rootCmd.Context(), "error running root command")
	}
}

func init() {
	rootCmd.AddCommand(NewServerCmd())
}
