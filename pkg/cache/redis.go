package cache

import (
	"context"
	"log"
	"os"

	"github.com/redis/go-redis/v9"
)

var (
	Rdb *redis.Client
	Ctx = context.Background()
)

// Init initializes Redis connection from environment variables
func Init() {
	Rdb = redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_HOST") + ":" + os.Getenv("REDIS_PORT"),
		Password: os.Getenv("REDIS_PASSWORD"), // "" if no password
		DB:       0,
	})

	// Test connection
	if err := Rdb.Ping(Ctx).Err(); err != nil {
		log.Fatalf("❌ Failed to connect Redis: %v", err)
	}

	log.Println("✅ Redis connected")
}
