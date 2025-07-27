package handlers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"

	"portfolio-backend/configs"
	"portfolio-backend/internal/middleware"
)

type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type AuthHandler struct {
	config *configs.Config
}

func NewAuthHandler(config *configs.Config) *AuthHandler {
	return &AuthHandler{
		config: config,
	}
}

func (h *AuthHandler) Login(c *gin.Context) {
	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// For demo purposes, using hardcoded admin credentials
	// In production, you should store these in the database with proper hashing
	adminUsername := "admin"
	adminPasswordHash := "$2a$10$92IXUNpkjO0rOQ5byMi.Ye4oKoEa3Ro9llC/.og/at2.uheWG/igi" // "password"

	if req.Username != adminUsername {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(adminPasswordHash), []byte(req.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	// Parse JWT expiry duration
	expiry, err := time.ParseDuration(h.config.JWTExpiry)
	if err != nil {
		expiry = 24 * time.Hour // Default to 24 hours
	}

	// Generate JWT token
	token, err := middleware.GenerateToken(req.Username, h.config.JWTSecret, expiry)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Login successful",
		"token":   token,
		"user": gin.H{
			"username": req.Username,
		},
	})
}
