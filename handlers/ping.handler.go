package handlers

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

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
	// Use Redis's DEL command to release the lock
	err := rdb.Del(ctx, key).Err()
	if err != nil {
		log.Printf("Failed to release lock for key %s: %v", key, err)
	}
}

func PingHandler(rdb *redis.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
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
}
