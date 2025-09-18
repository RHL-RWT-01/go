package handlers

import (
	"net/http"

	"github.com/RHL-RWT-01/go/auth"
	"github.com/RHL-RWT-01/go/models"
	"github.com/gin-gonic/gin"
)

// In-memory user storage (in production, this would be a database)
var users = make(map[string]*models.User)
var userIDCounter = 1

// Register handles user registration
func Register(c *gin.Context) {
	var req models.RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Check if user already exists
	if _, exists := users[req.Username]; exists {
		c.JSON(http.StatusConflict, gin.H{"error": "Username already exists"})
		return
	}

	// Create new user
	user := &models.User{
		ID:       userIDCounter,
		Username: req.Username,
		Email:    req.Email,
		Password: req.Password,
	}

	// Hash password
	if err := user.HashPassword(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}

	// Store user (in production, save to database)
	users[req.Username] = user
	userIDCounter++

	// Generate token
	token, err := auth.GenerateToken(user.ID, user.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	// Return response
	response := models.AuthResponse{
		Token: token,
		User: models.User{
			ID:       user.ID,
			Username: user.Username,
			Email:    user.Email,
		},
	}

	c.JSON(http.StatusCreated, response)
}

// Login handles user authentication
func Login(c *gin.Context) {
	var req models.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Find user
	user, exists := users[req.Username]
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	// Check password
	if !user.CheckPassword(req.Password) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	// Generate token
	token, err := auth.GenerateToken(user.ID, user.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	// Return response
	response := models.AuthResponse{
		Token: token,
		User: models.User{
			ID:       user.ID,
			Username: user.Username,
			Email:    user.Email,
		},
	}

	c.JSON(http.StatusOK, response)
}

// GetProfile returns the authenticated user's profile
func GetProfile(c *gin.Context) {
	_, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	username, exists := c.Get("username")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	// Find user by username (in production, you'd query by ID)
	user, exists := users[username.(string)]
	if !exists {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	// Return user profile
	profile := models.User{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
	}

	c.JSON(http.StatusOK, profile)
}

// ProtectedRoute is an example of a protected endpoint
func ProtectedRoute(c *gin.Context) {
	userID, _ := c.Get("userID")
	username, _ := c.Get("username")

	c.JSON(http.StatusOK, gin.H{
		"message":  "This is a protected route",
		"user_id":  userID,
		"username": username,
	})
}