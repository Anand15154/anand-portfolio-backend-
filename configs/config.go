package configs

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port            string
	GinMode         string
	MongoDBURI      string
	MongoDBDatabase string
	JWTSecret       string
	JWTExpiry       string
	AllowedOrigins  string
	SMTPHost        string
	SMTPPort        string
	SMTPUsername    string
	SMTPPassword    string
}

func LoadConfig() *Config {
	// Load .env file if it exists
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using environment variables")
	}

	config := &Config{
		Port:            getEnv("PORT", "8080"),
		GinMode:         getEnv("GIN_MODE", "debug"),
		MongoDBURI:      getEnv("MONGODB_URI", "mongodb://localhost:27017"),
		MongoDBDatabase: getEnv("MONGODB_DATABASE", "portfolio_db"),
		JWTSecret:       getEnv("JWT_SECRET", "your-super-secret-jwt-key-here"),
		JWTExpiry:       getEnv("JWT_EXPIRY", "24h"),
		AllowedOrigins:  getEnv("ALLOWED_ORIGINS", "http://localhost:5173,http://localhost:3000"),
		SMTPHost:        getEnv("SMTP_HOST", "smtp.gmail.com"),
		SMTPPort:        getEnv("SMTP_PORT", "587"),
		SMTPUsername:    getEnv("SMTP_USERNAME", ""),
		SMTPPassword:    getEnv("SMTP_PASSWORD", ""),
	}

	// Log configuration for debugging (without sensitive data)
	log.Printf("Configuration loaded:")
	log.Printf("  Port: %s", config.Port)
	log.Printf("  Gin Mode: %s", config.GinMode)
	log.Printf("  MongoDB Database: %s", config.MongoDBDatabase)
	log.Printf("  JWT Expiry: %s", config.JWTExpiry)
	log.Printf("  Allowed Origins: %s", config.AllowedOrigins)

	// Log MongoDB URI status (without exposing the actual URI)
	if config.MongoDBURI == "mongodb://localhost:27017" {
		log.Printf("  MongoDB URI: Using default localhost (MONGODB_URI not set)")
	} else {
		log.Printf("  MongoDB URI: Custom URI configured")
	}

	// Validate required configuration
	validateConfig(config)

	return config
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func validateConfig(config *Config) {
	// Check for critical configuration issues
	if config.MongoDBURI == "mongodb://localhost:27017" {
		log.Println("⚠️  WARNING: Using default MongoDB URI (localhost). Set MONGODB_URI environment variable for production.")
	}

	if config.JWTSecret == "your-super-secret-jwt-key-here" {
		log.Println("⚠️  WARNING: Using default JWT secret. Set JWT_SECRET environment variable for production.")
	}

	if config.GinMode == "debug" {
		log.Println("ℹ️  INFO: Running in debug mode. Set GIN_MODE=release for production.")
	}
}
