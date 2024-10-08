package config

import (
	"log"
	"strings"

	"github.com/spf13/viper"
)

// AppConfig holds all the configuration for the application
type AppConfig struct {
	ServerAddress string   `mapstructure:"server_address"`
	AllowedIPs    []string `mapstructure:"allowed_ips"`
	Environment   string   `mapstructure:"environment"`
	DatabaseURL   string   `mapstructure:"database_url"`
	OpenaiKey     string   `mapstructure:"openai_key"`
	ClaudeKey     string   `mapstructure:"claude_key"`
	ServiceType   string   `mapstructure:"service_type"`  // "mock" or "openai"
	DatabaseType  string   `mapstructure:"database_type"` // "sqlite" or "postgres"
	ContextString string   `mapstructure:"context_string"`
	ModelType     string   `mapstructure:"model_type"` // "fast", "normal", or "accurate"
}

// Config is the exported configuration object
var Config AppConfig

// Setup initializes the configuration
func Setup() {
	viper.SetConfigName("config") // Name of config file (without extension)
	viper.SetConfigType("yaml")   // Type of the config file
	viper.AddConfigPath("/app")   // Path to look for the config file in
	viper.AddConfigPath(".")      // Also look for config in the current directory
	viper.AutomaticEnv()          // Read in environment variables that match
	viper.SetDefault("allowed_ips", "127.0.0.1,::1")

	// Set defaults
	viper.SetDefault("server_address", ":8080")
	viper.SetDefault("environment", "development")
	viper.SetDefault("service_type", "mock") // Default to mock service
	viper.SetDefault("model_type", "normal")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}

	err := viper.Unmarshal(&Config)
	if err != nil {
		log.Fatalf("Unable to decode into struct, %s", err)
	}
	// Split allowed_ips into a slice
	Config.AllowedIPs = strings.Split(viper.GetString("allowed_ips"), ",")
}
