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

	// fields for service configuration
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

	// prompts for image analysis
	ClassifyImagePrompt   string `mapstructure:"classify_image_prompt"`
	FoodImagePrompt       string `mapstructure:"food_image_prompt"`
	NutritionLabelPrompt  string `mapstructure:"nutrition_label_prompt"`
	BarcodePrompt         string `mapstructure:"barcode_prompt"`
	DefaultImagePrompt    string `mapstructure:"default_image_prompt"`
	TextAnalysisPrompt    string `mapstructure:"text_analysis_prompt"`
	JSONFormatInstruction string `mapstructure:"json_format_instruction"`

	// field for mock service configuration
	MockServiceType string `mapstructure:"mock_service_type"`
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

	// Set defaults for new fields
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

	// Set defaults for new prompt fields
	viper.SetDefault("classify_image_prompt", "Classify this image as one of the following: food photo, nutrition label, barcode, or unknown. Respond with just the classification.")
	viper.SetDefault("food_image_prompt", "Analyze this food image and provide nutritional information.")
	viper.SetDefault("nutrition_label_prompt", "Extract and summarize the nutritional information from this nutrition label.")
	viper.SetDefault("barcode_prompt", "This is a barcode. If you can read it, provide the encoded information and any related nutritional data if available.")
	viper.SetDefault("default_image_prompt", "Analyze this image and provide any relevant nutritional information.")
	viper.SetDefault("text_analysis_prompt", "Analyze this food description and provide nutritional information.")
	viper.SetDefault("json_format_instruction", "Provide the response in JSON format with 'summary' and 'nutrition' fields.")

	// Set default for mock service type
	viper.SetDefault("mock_service_type", "default")

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
