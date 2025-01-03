package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var setConfigCmd = &cobra.Command{
	Use:   "set-config",
	Short: "Set a configuration value",
	Run: func(cmd *cobra.Command, args []string) {
		key, _ := cmd.Flags().GetString("key")
		value, _ := cmd.Flags().GetString("value")

		if key == "" || value == "" {
			log.Fatalln("Both key and value are required.")
		}

		viper.Set(key, value)

		if err := viper.WriteConfig(); err != nil {
			log.Fatalf("Error writing config file: %v\n", err)
		}

		fmt.Printf("Configuration %s set to %s\n", key, value)
	},
}

func init() {
	setConfigCmd.Flags().StringP("key", "k", "", "Configuration key")
	setConfigCmd.Flags().StringP("value", "v", "", "Configuration value")
	rootCmd.AddCommand(setConfigCmd)
}
