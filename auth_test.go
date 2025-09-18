package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/RHL-RWT-01/go/auth/handlers"
	"github.com/RHL-RWT-01/go/auth/middleware"
	"github.com/RHL-RWT-01/go/auth/models"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func setupTestRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	
	userStore := models.NewUserStore()
	authHandler := handlers.NewAuthHandler(userStore)
	
	r := gin.New()
	
	auth := r.Group("/auth")
	{
		auth.POST("/register", authHandler.Register)
		auth.POST("/login", authHandler.Login)
		
		protected := auth.Group("/")
		protected.Use(middleware.AuthMiddleware())
		{
			protected.GET("/profile", authHandler.GetProfile)
			protected.GET("/users", authHandler.GetUsers)
		}
	}
	
	return r
}

func TestUserRegistration(t *testing.T) {
	router := setupTestRouter()
	
	registerData := models.RegisterRequest{
		Username: "testuser",
		Email:    "test@example.com",
		Password: "password123",
	}
	
	jsonData, _ := json.Marshal(registerData)
	req, _ := http.NewRequest("POST", "/auth/register", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")
	
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	
	assert.Equal(t, http.StatusCreated, w.Code)
	
	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, "User registered successfully", response["message"])
	assert.NotEmpty(t, response["token"])
}

func TestUserLogin(t *testing.T) {
	router := setupTestRouter()
	
	// First register a user
	registerData := models.RegisterRequest{
		Username: "testuser",
		Email:    "test@example.com",
		Password: "password123",
	}
	
	jsonData, _ := json.Marshal(registerData)
	req, _ := http.NewRequest("POST", "/auth/register", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")
	
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	
	// Now test login
	loginData := models.LoginRequest{
		Username: "testuser",
		Password: "password123",
	}
	
	jsonData, _ = json.Marshal(loginData)
	req, _ = http.NewRequest("POST", "/auth/login", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")
	
	w = httptest.NewRecorder()
	router.ServeHTTP(w, req)
	
	assert.Equal(t, http.StatusOK, w.Code)
	
	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, "Login successful", response["message"])
	assert.NotEmpty(t, response["token"])
}

func TestProtectedEndpointWithoutAuth(t *testing.T) {
	router := setupTestRouter()
	
	req, _ := http.NewRequest("GET", "/auth/profile", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	
	assert.Equal(t, http.StatusUnauthorized, w.Code)
}

func TestDuplicateUserRegistration(t *testing.T) {
	router := setupTestRouter()
	
	registerData := models.RegisterRequest{
		Username: "testuser",
		Email:    "test@example.com",
		Password: "password123",
	}
	
	// Register first user
	jsonData, _ := json.Marshal(registerData)
	req, _ := http.NewRequest("POST", "/auth/register", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")
	
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusCreated, w.Code)
	
	// Try to register same user again
	req, _ = http.NewRequest("POST", "/auth/register", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")
	
	w = httptest.NewRecorder()
	router.ServeHTTP(w, req)
	
	assert.Equal(t, http.StatusConflict, w.Code)
	
	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, "username already exists", response["error"])
}