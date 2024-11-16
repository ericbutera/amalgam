package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "amalgam-cli",
	Short: "amalgam-cli",
	Long:  "amalgam-cli",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
	os.Exit(0)
}

func init() {
	feedCmd := NewFeedCmd()
	feedCmd.AddCommand(NewFeedListCmd())
	feedCmd.AddCommand(NewFeedAddCmd())
	rootCmd.AddCommand(feedCmd)
}
