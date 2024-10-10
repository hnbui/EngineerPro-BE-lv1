package main

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"

	"engineerpro/configs"
	"engineerpro/handlers"
)

func main() {
	// Set up the Gin router
	router := gin.Default()

	// Initialize Redis client
	rdb := configs.InitRedis()

	// Create Redis session storage
	store := configs.CreateRedisStore()
	router.Use(sessions.Sessions("loginsession", store))

	// Define routes
	router.GET("/ping", handlers.PingHandler(rdb))
	router.POST("/login", handlers.LoginHandler)

	// Start the Gin server
	router.Run(":8080")
}
