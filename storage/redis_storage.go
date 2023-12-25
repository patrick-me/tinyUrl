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

func (s *RedisStorage) Get(short string) (string, error) {
	if val, err := s.Client.Get(s.Context, short).Result(); errors.Is(err, redis.Nil) {
		return val, nil
	}

	return "", fmt.Errorf("URL is not found: %s", short)
}
