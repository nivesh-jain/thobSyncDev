package main

import (
	"fmt"
	"log"

	"github.com/nivesh-jain/thobSyncDev.git/cmd"

	"github.com/spf13/viper"
)

func initConfig() {
	viper.SetConfigName("config") // Name of config file (without extension)
	viper.SetConfigType("yaml")   // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath(".")      // Look for config in the working directory

	// Set default values
	viper.SetDefault("minio.endpoint", "localhost:9000")
	viper.SetDefault("minio.accessKeyID", "Gx0S3h31P8SfmOWhm3Tg")
	viper.SetDefault("minio.secretAccessKey", "XAqfnX6Q77PhtEUhyjziZj8bsPpz9PoSLtgSh1yY")
	viper.SetDefault("minio.useSSL", false)

	// Read in environment variables that match
	viper.AutomaticEnv()

	// If a config file is found, read it in
	if err := viper.ReadInConfig(); err != nil {
		log.Printf("Error reading config file: %v\n", err)
	} else {
		fmt.Printf("Using config file: %s\n", viper.ConfigFileUsed())
	}
}

func main() {
	initConfig()
	cmd.Execute()
}
