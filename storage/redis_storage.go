package storage

import (
	"context"
	"errors"
	"fmt"
	"github.com/redis/go-redis/v9"
	"time"
)

type RedisStorage struct {
	Client     *redis.Client
	Context    context.Context
	Expiration time.Duration
}

func (s *RedisStorage) Contains(short string) bool {
	_, err := s.Client.Get(s.Context, short).Result()
	return !errors.Is(err, redis.Nil)
}

func (s *RedisStorage) Save(short, origin string) {
	s.Client.Set(s.Context, short, origin, s.Expiration)
}

func (s *RedisStorage) Get(short string) (val string, err error) {
	val, err = s.Client.Get(s.Context, short).Result()

	if !errors.Is(err, redis.Nil) {
		return val, nil
	}

	fmt.Printf("URL is not found: %s, err: %e\n", short, err)

	return "", fmt.Errorf("URL is not found: %s", short)
}
