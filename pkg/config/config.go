package config

import (
	"log"

	"github.com/spf13/viper"
)

// AppConfig holds all the configuration for the application
type AppConfig struct {
	ServerAddress string
	Environment   string
	DatabaseURL   string
}

// Config is the exported configuration object
var Config AppConfig

// Setup initializes the configuration
func Setup() {
	viper.SetConfigName("config") // Name of config file (without extension)
	viper.SetConfigType("yaml")   // Type of the config file
	viper.AddConfigPath("/app")   // Path to look for the config file in
	viper.AutomaticEnv()          // Read in environment variables that match

	// Set defaults
	viper.SetDefault("server_address", ":8080")
	viper.SetDefault("environment", "development")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}

	err := viper.Unmarshal(&Config)
	if err != nil {
		log.Fatalf("Unable to decode into struct, %s", err)
	}
}
