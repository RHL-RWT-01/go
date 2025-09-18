package config

import (
	"os"
)

// Config holds the application configuration
type Config struct {
	Port      string
	JWTSecret string
}

// Load loads the configuration from environment variables
func Load() *Config {
	return &Config{
		Port:      getEnv("PORT", "8080"),
		JWTSecret: getEnv("JWT_SECRET", "your-secret-key"),
	}
}

// getEnv gets an environment variable with a fallback value
func getEnv(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}