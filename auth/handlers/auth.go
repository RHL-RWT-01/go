package handlers

import (
	"net/http"

	"github.com/RHL-RWT-01/go/auth/models"
	"github.com/RHL-RWT-01/go/auth/utils"
	"github.com/gin-gonic/gin"
)

// AuthHandler handles authentication-related requests
type AuthHandler struct {
	userStore *models.UserStore
}

// NewAuthHandler creates a new authentication handler
func NewAuthHandler(userStore *models.UserStore) *AuthHandler {
	return &AuthHandler{
		userStore: userStore,
	}
}

// Register handles user registration
func (h *AuthHandler) Register(c *gin.Context) {
	var req models.RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request data: " + err.Error(),
		})
		return
	}

	// Hash the password
	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to hash password",
		})
		return
	}

	// Create user
	user, err := h.userStore.CreateUser(req.Username, req.Email, hashedPassword)
	if err != nil {
		c.JSON(http.StatusConflict, gin.H{
			"error": err.Error(),
		})
		return
	}

	// Generate token
	token, err := utils.GenerateToken(user.ID, user.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to generate token",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "User registered successfully",
		"token":   token,
		"user":    user.ToUserResponse(),
	})
}

// Login handles user authentication
func (h *AuthHandler) Login(c *gin.Context) {
	var req models.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request data: " + err.Error(),
		})
		return
	}

	// Get user by username
	user, err := h.userStore.GetUserByUsername(req.Username)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Invalid credentials",
		})
		return
	}

	// Check password
	if err := utils.CheckPassword(req.Password, user.Password); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Invalid credentials",
		})
		return
	}

	// Generate token
	token, err := utils.GenerateToken(user.ID, user.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to generate token",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Login successful",
		"token":   token,
		"user":    user.ToUserResponse(),
	})
}

// GetProfile returns the current user's profile
func (h *AuthHandler) GetProfile(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "User not authenticated",
		})
		return
	}

	user, err := h.userStore.GetUserByID(userID.(int))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "User not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"user": user.ToUserResponse(),
	})
}

// GetUsers returns all users (protected endpoint for demonstration)
func (h *AuthHandler) GetUsers(c *gin.Context) {
	users := h.userStore.GetAllUsers()
	userResponses := make([]models.UserResponse, len(users))
	
	for i, user := range users {
		userResponses[i] = user.ToUserResponse()
	}

	c.JSON(http.StatusOK, gin.H{
		"users": userResponses,
	})
}