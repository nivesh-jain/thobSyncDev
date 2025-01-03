package config

import (
	"log"

	"github.com/spf13/viper"
)

func InitConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	// Set default values
	viper.SetDefault("minio.endpoint", "localhost:9000")
	viper.SetDefault("minio.accessKeyID", "Gx0S3h31P8SfmOWhm3Tg")
	viper.SetDefault("minio.secretAccessKey", "XAqfnX6Q77PhtEUhyjziZj8bsPpz9PoSLtgSh1yY")
	viper.SetDefault("minio.useSSL", false)

	// Read the config file
	if err := viper.ReadInConfig(); err != nil {
		log.Printf("No config file found: %v\n", err)
	}
}
