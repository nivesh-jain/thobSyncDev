package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "thobSyncDev",
	Short: "A CLI tool for 3D artists",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(listBucketsCmd)
	rootCmd.AddCommand(uploadFileCmd)
	rootCmd.AddCommand(downloadFileCmd)
	rootCmd.AddCommand(deleteFileCmd)
	rootCmd.AddCommand(listFilesCmd)
}
