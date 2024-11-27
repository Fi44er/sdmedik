package redis

import (
	"context"
	"fmt"
	"os"

	"github.com/redis/go-redis/v9"
)

var ctx context.Context

type Cache struct {
	RedisClient *redis.Client
}

var RedisCache Cache

func Connect() (Cache, error) {
	ctx = context.Background()

	redisClient := redis.NewClient(&redis.Options{
		Addr: os.Getenv("REDIS_URL"),
	})

	if _, err := redisClient.Ping(ctx).Result(); err != nil {
		return Cache{}, err
	}

	RedisCache = Cache{
		RedisClient: redisClient,
	}

	fmt.Println("✅ Redis client connected successfully...")

	return RedisCache, nil
}
