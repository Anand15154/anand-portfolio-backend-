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

	return &Config{
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
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
