package storage

import "time"

type Storage interface {
	Contains(short string) bool
	Save(short, origin string, expirationInHours time.Duration)
	Get(short string) (string, error)
}
