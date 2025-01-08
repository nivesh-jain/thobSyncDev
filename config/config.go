package config

import (
	"log"

	"github.com/spf13/viper"
)

// InitConfig initializes the configuration system.
func InitConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	// Set default values for MinIO
	viper.SetDefault("minio.endpoint", "localhost:9000")
	viper.SetDefault("minio.accessKeyID", "Gx0S3h31P8SfmOWhm3Tg")
	viper.SetDefault("minio.secretAccessKey", "XAqfnX6Q77PhtEUhyjziZj8bsPpz9PoSLtgSh1yY")
	viper.SetDefault("minio.useSSL", false)

	// Set default user structure (empty on first run)
	viper.SetDefault("user.username", "")
	viper.SetDefault("user.role", "")
	viper.SetDefault("user.token", "")

	// Attempt to read the config file
	if err := viper.ReadInConfig(); err != nil {
		log.Printf("No config file found, using defaults: %v\n", err)
	}
}

// GenerateConfigFile writes user-specific details to the config file.
func GenerateConfigFile(username, role, token string) {
	viper.Set("user.username", username)
	viper.Set("user.role", role)
	viper.Set("user.token", token)

	// Save the updated configuration to `config.yaml`
	if err := viper.WriteConfigAs("config.yaml"); err != nil {
		log.Fatalf("Failed to write config file: %v", err)
	}
	log.Println("Config file generated successfully.")
}

// GetCurrentUser retrieves the current user's details from the config file.
func GetCurrentUser() (string, string) {
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Failed to read config file: %v", err)
	}

	username := viper.GetString("user.username")
	role := viper.GetString("user.role")
	return username, role
}

// GetMinIOConfig retrieves MinIO configuration details from the config file.
func GetMinIOConfig() (endpoint, accessKeyID, secretAccessKey string, useSSL bool) {
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Failed to read config file: %v", err)
	}

	endpoint = viper.GetString("minio.endpoint")
	accessKeyID = viper.GetString("minio.accessKeyID")
	secretAccessKey = viper.GetString("minio.secretAccessKey")
	useSSL = viper.GetBool("minio.useSSL")

	return
}
