package db

import (
	"github.com/redis/go-redis/v9"
	"tinyUrl/utils"
)

func CreateRedisClient() *redis.Client {
	client := redis.NewClient(
		&redis.Options{
			Addr:     utils.GetEnv("REDIS_ADDR", ":6379"),
			Password: utils.GetEnv("REDIS_PASSWORD", ""),
		})

	return client
}
