package configs

import (
	"context"
	"log"

	redisStore "github.com/gin-contrib/sessions/redis"
	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

func InitRedis() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		log.Fatalf("Failed to connect to Redis: %v", err)
	}

	log.Println("Connected to Redis server")

	return rdb
}

func CreateRedisStore() redisStore.Store {
	store, err := redisStore.NewStore(10, "tcp", "localhost:6379", "", []byte("test"))
	if err != nil {
		log.Fatalf("Failed to create Redis store: %v", err)
	}

	log.Println("Created Redis store")

	return store
}
