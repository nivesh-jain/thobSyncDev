package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var showConfigCmd = &cobra.Command{
	Use:   "show-config",
	Short: "Display the current configuration",
	Run: func(cmd *cobra.Command, args []string) {
		settings := viper.AllSettings()
		for key, value := range settings {
			fmt.Printf("%s: %v\n", key, value)
		}
	},
}

func init() {
	rootCmd.AddCommand(showConfigCmd)
}
