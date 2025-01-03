package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "3d-cli-tool",
	Short: "A CLI tool for 3D artists",
	Long:  "A command-line tool providing collaboration features for 3D artists.",
}

// Execute runs the root command
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func init() {
	// Add subcommands here
	rootCmd.AddCommand(listBucketsCmd)
}
