package main

import (
	"log"
	"net/http"

	"github.com/RHL-RWT-01/go/handlers"
	"github.com/RHL-RWT-01/go/middleware"
	"github.com/gin-gonic/gin"
)

func main() {
	// Create Gin router
	r := gin.Default()

	// Add CORS middleware for development
	r.Use(func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}

		c.Next()
	})

	// Health check endpoint
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "healthy",
			"message": "JWT Auth API is running",
		})
	})

	// Public routes
	auth := r.Group("/auth")
	{
		auth.POST("/register", handlers.Register)
		auth.POST("/login", handlers.Login)
	}

	// Protected routes
	api := r.Group("/api")
	api.Use(middleware.JWTAuthMiddleware())
	{
		api.GET("/profile", handlers.GetProfile)
		api.GET("/protected", handlers.ProtectedRoute)
	}

	// API documentation endpoint
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "JWT Authentication API",
			"version": "1.0.0",
			"endpoints": gin.H{
				"public": gin.H{
					"POST /auth/register": "Register a new user",
					"POST /auth/login":    "Login with username and password",
					"GET /health":         "Health check",
				},
				"protected": gin.H{
					"GET /api/profile":   "Get user profile (requires Authorization header)",
					"GET /api/protected": "Example protected route (requires Authorization header)",
				},
			},
			"auth_header_format": "Authorization: Bearer <jwt_token>",
		})
	})

	// Start server
	log.Println("Server starting on :8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}