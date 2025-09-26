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

func InitRedis() {
	Rdb = redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_HOST") + ":" + os.Getenv("REDIS_PORT"),
		Password: os.Getenv("REDIS_PASSWORD"), // "" if no password
		DB:       0,
	})

	// Test connection
	_, err := Rdb.Ping(Ctx).Result()
	if err != nil {
		log.Fatal("❌ Failed to connect to Redis:", err)
	}
	log.Println("✅ Redis connected")
}
