package main

import (
	"fmt"
	"os"

	"portfolio-backend/configs"
)

func main() {
	fmt.Println("🔍 Environment Variable Checker")
	fmt.Println("================================")

	// Load configuration
	config := configs.LoadConfig()

	fmt.Println("\n📋 Configuration Summary:")
	fmt.Printf("  Port: %s\n", config.Port)
	fmt.Printf("  Gin Mode: %s\n", config.GinMode)
	fmt.Printf("  MongoDB Database: %s\n", config.MongoDBDatabase)
	fmt.Printf("  JWT Expiry: %s\n", config.JWTExpiry)
	fmt.Printf("  Allowed Origins: %s\n", config.AllowedOrigins)

	// Check MongoDB URI
	if config.MongoDBURI == "mongodb://localhost:27017" {
		fmt.Println("\n❌ MONGODB_URI: Not set (using default localhost)")
		fmt.Println("   Set MONGODB_URI environment variable to your MongoDB Atlas connection string")
	} else {
		fmt.Println("\n✅ MONGODB_URI: Configured")
	}

	// Check JWT Secret
	if config.JWTSecret == "your-super-secret-jwt-key-here" {
		fmt.Println("❌ JWT_SECRET: Not set (using default)")
		fmt.Println("   Set JWT_SECRET environment variable for production")
	} else {
		fmt.Println("✅ JWT_SECRET: Configured")
	}

	// Check Gin Mode
	if config.GinMode == "debug" {
		fmt.Println("ℹ️  GIN_MODE: Debug mode (set to 'release' for production)")
	} else {
		fmt.Println("✅ GIN_MODE: Production mode")
	}

	fmt.Println("\n🚀 For Railway deployment, set these environment variables:")
	fmt.Println("   MONGODB_URI=mongodb+srv://your-username:your-password@your-cluster.mongodb.net/")
	fmt.Println("   MONGODB_DATABASE=portfolio_db")
	fmt.Println("   JWT_SECRET=your-super-secret-jwt-key-here")
	fmt.Println("   GIN_MODE=release")
	fmt.Println("   ALLOWED_ORIGINS=https://your-frontend-domain.com")

	// Check if running in Railway
	if os.Getenv("RAILWAY_ENVIRONMENT") != "" {
		fmt.Println("\n🚂 Running in Railway environment")
	} else {
		fmt.Println("\n💻 Running in local environment")
	}
}
