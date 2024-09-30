package configs

import (
	"crypto/rand"
	"encoding/hex"

	"github.com/gin-gonic/gin"
)

// Helper function to create a session ID
func GenerateSessionID() (string, error) {
	bytes := make([]byte, 16) // 16 bytes = 128-bit token
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}

	return hex.EncodeToString(bytes), nil
}

// Middleware to validate session
func SessionAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get the session ID from the request header

	}
}
