package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"portfolio-backend/configs"
	"portfolio-backend/internal/database"
	"portfolio-backend/internal/handlers"
	"portfolio-backend/internal/middleware"
	"portfolio-backend/internal/services"
)

func main() {
	// Load configuration
	config := configs.LoadConfig()

	// Set Gin mode
	gin.SetMode(config.GinMode)

	// Initialize database
	db, err := database.NewMongoDB(config.MongoDBURI, config.MongoDBDatabase)
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	defer db.Close()

	// Initialize email service
	emailService := services.NewEmailService(
		config.SMTPHost,
		config.SMTPPort,
		config.SMTPUsername,
		config.SMTPPassword,
	)

	// Initialize services
	contactService := services.NewContactService(db, emailService)
	projectService := services.NewProjectService(db)

	// Initialize handlers
	contactHandler := handlers.NewContactHandler(contactService)
	projectHandler := handlers.NewProjectHandler(projectService)
	authHandler := handlers.NewAuthHandler(config)

	// Initialize rate limiter
	rateLimiter := middleware.NewRateLimiter(10, time.Minute) // 10 requests per minute

	// Initialize router
	router := gin.New()

	// Add middleware
	router.Use(middleware.LoggingMiddleware())
	router.Use(middleware.CORSMiddleware(config.AllowedOrigins))
	router.Use(gin.Recovery())

	// Health check endpoint
	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status":  "ok",
			"message": "Portfolio backend is running",
			"time":    time.Now().Format(time.RFC3339),
		})
	})

	// API routes
	api := router.Group("/api/v1")
	{
		// Auth routes
		auth := api.Group("/auth")
		{
			auth.POST("/login", authHandler.Login)
		}

		// Contact routes
		contacts := api.Group("/contacts")
		{
			contacts.POST("/", rateLimiter.RateLimitMiddleware(), contactHandler.CreateContact)
			contacts.GET("/", middleware.AuthMiddleware(config.JWTSecret), contactHandler.GetAllContacts)
			contacts.GET("/:id", middleware.AuthMiddleware(config.JWTSecret), contactHandler.GetContactByID)
			contacts.PUT("/:id/read", middleware.AuthMiddleware(config.JWTSecret), contactHandler.MarkAsRead)
			contacts.DELETE("/:id", middleware.AuthMiddleware(config.JWTSecret), contactHandler.DeleteContact)
		}

		// Project routes
		projects := api.Group("/projects")
		{
			projects.POST("/", middleware.AuthMiddleware(config.JWTSecret), projectHandler.CreateProject)
			projects.GET("/", projectHandler.GetAllProjects)
			projects.GET("/featured", projectHandler.GetFeaturedProjects)
			projects.GET("/:id", projectHandler.GetProjectByID)
			projects.PUT("/:id", middleware.AuthMiddleware(config.JWTSecret), projectHandler.UpdateProject)
			projects.DELETE("/:id", middleware.AuthMiddleware(config.JWTSecret), projectHandler.DeleteProject)
		}
	}

	// Start server
	log.Printf("Server starting on port %s", config.Port)
	if err := router.Run(":" + config.Port); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
