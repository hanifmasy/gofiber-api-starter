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
	// Get Redis configuration from environment variables with defaults
	redisHost := os.Getenv("REDIS_HOST")
	if redisHost == "" {
		redisHost = "localhost"
	}

	redisPort := os.Getenv("REDIS_PORT")
	if redisPort == "" {
		redisPort = "6379"
	}

	redisPassword := os.Getenv("REDIS_PASSWORD") // "" if no password

	Rdb = redis.NewClient(&redis.Options{
		Addr:     redisHost + ":" + redisPort,
		Password: redisPassword, // "" if no password
		DB:       0,
	})

	// Test connection - but don't kill the app if it fails
	if err := Rdb.Ping(Ctx).Err(); err != nil {
		log.Printf("❌ Failed to connect Redis: %v", err)
		log.Println("⚠️  Application will continue without Redis")
		Rdb = nil // Set to nil so cache functions can check availability
		return
	}

	log.Println("✅ Redis connected")
}

// IsConnected checks if Redis connection is available
func IsConnected() bool {
	return Rdb != nil
}
