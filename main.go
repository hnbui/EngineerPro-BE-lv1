package main

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"engineerpro/configs"
)

func main() {
	// Initialize Redis connection
	configs.InitRedis()

	// Set up the Gin router
	router := gin.Default()

	// Define routes
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	// Start the Gin server
	router.Run(":8080")
}
