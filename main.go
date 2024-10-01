package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"engineerpro/configs"
)

func main() {
	// Initialize Redis connection
	err := configs.InitRedis()
	if err != nil {
		panic(fmt.Sprintf("Couldn't connect to Redis: %v", err))
	}

	// Initialize MySQL database connection
	err = configs.InitDB()
	if err != nil {
		panic(fmt.Sprintf("Couldn't connect to MySQL: %v", err))
	}

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
