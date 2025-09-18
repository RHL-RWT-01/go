package main

import (
	"log"
	"net/http"

	"github.com/RHL-RWT-01/go/auth/handlers"
	"github.com/RHL-RWT-01/go/auth/middleware"
	"github.com/RHL-RWT-01/go/auth/models"
	"github.com/RHL-RWT-01/go/config"
	"github.com/gin-gonic/gin"
)

func main() {
	// Load configuration
	cfg := config.Load()

	// Initialize user store
	userStore := models.NewUserStore()

	// Initialize handlers
	authHandler := handlers.NewAuthHandler(userStore)

	// Setup Gin router
	r := gin.Default()

	// Add CORS middleware for development
	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	})

	// Public routes
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Welcome to the Authentication API",
			"version": "1.0.0",
			"endpoints": gin.H{
				"register": "POST /auth/register",
				"login":    "POST /auth/login",
				"profile":  "GET /auth/profile (requires auth)",
				"users":    "GET /auth/users (requires auth)",
			},
		})
	})

	// Health check endpoint
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "healthy",
		})
	})

	// Authentication routes
	auth := r.Group("/auth")
	{
		// Public auth routes
		auth.POST("/register", authHandler.Register)
		auth.POST("/login", authHandler.Login)

		// Protected auth routes
		protected := auth.Group("/")
		protected.Use(middleware.AuthMiddleware())
		{
			protected.GET("/profile", authHandler.GetProfile)
			protected.GET("/users", authHandler.GetUsers)
		}
	}

	// Start server
	log.Printf("Server starting on port %s", cfg.Port)
	log.Printf("API Documentation:")
	log.Printf("  POST /auth/register - Register a new user")
	log.Printf("  POST /auth/login    - Login with username/password")
	log.Printf("  GET  /auth/profile  - Get user profile (requires Bearer token)")
	log.Printf("  GET  /auth/users    - Get all users (requires Bearer token)")
	log.Printf("  GET  /health        - Health check")
	
	if err := r.Run(":" + cfg.Port); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}