package configs

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
)

// Create a Redis client
var ctx = context.Background()
var rdb *redis.Client

// Initialize Redis connection
func InitRedis() {
	rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	// Test connection
	pong, err := rdb.Ping(ctx).Result()
	if err != nil {
		panic(fmt.Sprintf("Couldn't connect to Redis: %v", err))
	}
	fmt.Println("Redis connected:", pong)
}
