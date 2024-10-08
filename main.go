/*
	1. Stores session for users logging in.
	2. Limits access to /ping to allow only one user at a time (with a 5-second delay).
	3. Counts requests to /ping since the server started.
	4. Implements rate limiting of two requests to /ping per user in 60 seconds.
	5. Returns the top 10 users who called /ping the most.
	6. Uses HyperLogLog to track unique users calling /ping.
*/

/*
When a user logs in, a session is created and stored with a unique session ID, which is sent to the user's browser as a cookie.
Each time the user makes a request, their browser sends this session cookie to the server.
This cookie allows the server to know which session belongs to which user.
*/

package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
)

var (
	ctx       = context.Background()
	rdb       *redis.Client
	mu        sync.Mutex
	pingCount int
)

type Credentials struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func InitRedis() *redis.Client {
	rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	// Test connection
	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connect to Redis server")

	return rdb
}

func LoginHandler(c *gin.Context) {
	session := sessions.Default(c)
	var cred Credentials

	if err := c.ShouldBindJSON(&cred); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid crendential"})
		return
	}

	// Mock credentials
	if strings.Contains(cred.Username, "test") && cred.Password == "123456" {
		session.Set("Username", cred.Username)
		session.Save()
		c.JSON(http.StatusOK, gin.H{"message": "Login successful"})
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Wrong username or password"})
	}
}

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		sessions := sessions.Default(c)
		username := sessions.Get("Username")
		if username == nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized crendential"})
			return
		}
		c.Next()
	}
}

func AcquireLock(rdb *redis.Client, key string, expiration time.Duration) bool {
	// Use Redis's SetNX command to acquire a lock
	ok, err := rdb.SetNX(ctx, key, "locked", expiration).Result()
	if err != nil {
		log.Printf("Failed to acquire lock for key %s: %v", key, err)
		return false
	}
	return ok
}

func ReleaseLock(rdb *redis.Client, key string) {
	_, err := rdb.Del(ctx, key).Result()
	if err != nil {
		log.Printf("Failed to release lock for key %s: %v", key, err)
	}
}

func PingHandler(c *gin.Context) {
	lockKey := "ping"
	lockExpiration := 5 * time.Second

	// Try to acquire lock
	if !AcquireLock(rdb, lockKey, lockExpiration) {
		c.JSON(http.StatusTooManyRequests, gin.H{"error": "Another user is already accessing this API. Please try again later."})
		return
	}

	time.Sleep(5 * time.Second)

	ReleaseLock(rdb, lockKey)

	c.JSON(http.StatusOK, gin.H{"message": "Ping successful!"})
}

func main() {
	// Set up the Gin router
	router := gin.Default()

	// // Initialize Redis client
	// rdb = InitRedis()

	// // Initialize Redis storage
	// store, err := redisStore.NewStore(10, "tcp", "localhost:6379", "", []byte("test"))
	// if err != nil {
	// 	log.Fatalf("Failed to create Redis store %v", err)
	// }

	// router.Use(sessions.Sessions("loginsession", store))

	// Use a cookie-based session store instead of Redis for this test
	store := cookie.NewStore([]byte("secret"))
	router.Use(sessions.Sessions("loginsession", store))

	// Define routes
	router.POST("/login", LoginHandler)

	router.GET("/ping", PingHandler)

	// authorized := router.Group("/")
	// // authorized.Use(AuthMiddleware())
	// {
	// 	authorized.GET("/ping", PingHandler)
	// }

	// Start the Gin server
	router.Run(":8080")
}
