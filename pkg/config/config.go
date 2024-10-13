package config

import (
	"fmt"
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

	// Fields for service configuration
	ImageClassifierService       string `mapstructure:"image_classifier_service"`
	OpenAIModelForClassification string `mapstructure:"openai_model_for_classification"`
	ClaudeModelForClassification string `mapstructure:"claude_model_for_classification"`

	BarcodeAnalyzerService        string `mapstructure:"barcode_analyzer_service"`
	FoodImageAnalyzerService      string `mapstructure:"food_image_analyzer_service"`
	NutritionLabelAnalyzerService string `mapstructure:"nutrition_label_analyzer_service"`
	TextAnalyzerService           string `mapstructure:"text_analyzer_service"`
	DefaultAnalyzerService        string `mapstructure:"default_analyzer_service"`

	OpenAIModelForAnalysis string `mapstructure:"openai_model_for_analysis"`
	ClaudeModelForAnalysis string `mapstructure:"claude_model_for_analysis"`

	// Field for mock service configuration
	MockServiceType string `mapstructure:"mock_service_type"`

	// New field for prompt configurations
	Prompts map[string]LLMPrompts `mapstructure:"prompts"`
}

// LLMPrompts holds the prompts for a specific LLM
type LLMPrompts struct {
	File    string            `mapstructure:"file"`
	Prompts map[string]string `mapstructure:"-"` // We'll load these dynamically
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

	// Set defaults
	setDefaults()

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}

	err := viper.Unmarshal(&Config)
	if err != nil {
		log.Fatalf("Unable to decode into struct, %s", err)
	}

	// Split allowed_ips into a slice
	Config.AllowedIPs = strings.Split(viper.GetString("allowed_ips"), ",")

	// Load prompts for each LLM
	for llm, promptConfig := range Config.Prompts {
		if err := loadPrompts(llm, promptConfig.File); err != nil {
			log.Fatalf("Error loading prompts for %s: %s", llm, err)
		}
	}
}

// setDefaults sets the default values for configuration
func setDefaults() {
	viper.SetDefault("allowed_ips", "127.0.0.1,::1")
	viper.SetDefault("server_address", ":8080")
	viper.SetDefault("environment", "development")
	viper.SetDefault("service_type", "mock")
	viper.SetDefault("model_type", "normal")
	viper.SetDefault("image_classifier_service", "openai")
	viper.SetDefault("openai_model_for_classification", "gpt-4-vision-preview")
	viper.SetDefault("claude_model_for_classification", "claude-3-opus-20240229")
	viper.SetDefault("barcode_analyzer_service", "openai")
	viper.SetDefault("food_image_analyzer_service", "openai")
	viper.SetDefault("nutrition_label_analyzer_service", "openai")
	viper.SetDefault("text_analyzer_service", "openai")
	viper.SetDefault("default_analyzer_service", "openai")
	viper.SetDefault("openai_model_for_analysis", "gpt-4-vision-preview")
	viper.SetDefault("claude_model_for_analysis", "claude-3-opus-20240229")
	viper.SetDefault("mock_service_type", "default")

	// Set default for prompts
	viper.SetDefault("prompts", map[string]LLMPrompts{
		"openai": {File: "config/prompts/openai.yaml"},
		"claude": {File: "config/prompts/claude.yaml"},
		"llama":  {File: "config/prompts/llama.yaml"},
	})
}

// loadPrompts loads prompts from a file for a specific LLM
func loadPrompts(llm, file string) error {
	v := viper.New()
	v.SetConfigFile(file)
	if err := v.ReadInConfig(); err != nil {
		return fmt.Errorf("error reading prompt file: %w", err)
	}

	prompts := make(map[string]string)
	if err := v.Unmarshal(&prompts); err != nil {
		return fmt.Errorf("unable to decode prompts: %w", err)
	}

	Config.Prompts[llm] = LLMPrompts{
		File:    file,
		Prompts: prompts,
	}

	return nil
}

// GetPrompt retrieves a specific prompt for an LLM
func (c *AppConfig) GetPrompt(llm, promptName string) string {
	if llmPrompts, ok := c.Prompts[llm]; ok {
		if prompt, ok := llmPrompts.Prompts[promptName]; ok {
			return prompt
		}
	}
	return "" // or return a default prompt
}
