package db

import (
	"github.com/patrick-me/tinyUrl/utils"
	"github.com/redis/go-redis/v9"
)

func CreateRedisClient() *redis.Client {
	client := redis.NewClient(
		&redis.Options{
			Addr:     utils.GetEnv("REDIS_ADDR", ":6379"),
			Password: utils.GetEnv("REDIS_PASSWORD", ""),
		})

	return client
}
