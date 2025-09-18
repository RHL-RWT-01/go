package models

import (
	"time"
)

// User represents a user in the system
type User struct {
	ID        int       `json:"id"`
	Username  string    `json:"username" binding:"required,min=3,max=20"`
	Email     string    `json:"email" binding:"required,email"`
	Password  string    `json:"password,omitempty" binding:"required,min=6"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// UserResponse represents the user data returned in responses (without password)
type UserResponse struct {
	ID        int       `json:"id"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// LoginRequest represents the login request payload
type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// RegisterRequest represents the registration request payload
type RegisterRequest struct {
	Username string `json:"username" binding:"required,min=3,max=20"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

// ToUserResponse converts User to UserResponse (removes password)
func (u *User) ToUserResponse() UserResponse {
	return UserResponse{
		ID:        u.ID,
		Username:  u.Username,
		Email:     u.Email,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}
}