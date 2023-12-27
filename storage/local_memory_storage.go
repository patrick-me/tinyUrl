package storage

import (
	"fmt"
	"sync"
	"time"
)

type LocalMemoryStorage struct {
	sync.RWMutex
	Store map[string]string
}

func (s *LocalMemoryStorage) Contains(short string) bool {
	s.RLock()
	defer s.RUnlock()

	_, contains := s.Store[short]
	return contains
}

func (s *LocalMemoryStorage) Save(short, origin string, _ time.Duration) {
	s.Lock()
	defer s.Unlock()

	s.Store[short] = origin
}

func (s *LocalMemoryStorage) Get(short string) (string, error) {
	s.RLock()
	defer s.RUnlock()

	if val, exists := s.Store[short]; exists {
		return val, nil
	}

	return "", fmt.Errorf("URL is not found: %s", short)
}
