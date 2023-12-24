package storage

import (
	"fmt"
	"sync"
)

type LocalMemoryStorage struct {
	mutex sync.Mutex
	Store map[string]string
}

func (s *LocalMemoryStorage) Contains(short string) bool {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	_, contains := s.Store[short]
	return contains
}

func (s *LocalMemoryStorage) Save(short, origin string) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	s.Store[short] = origin
}

func (s *LocalMemoryStorage) Get(short string) (string, error) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	if val, exists := s.Store[short]; exists {
		return val, nil
	}

	return "", fmt.Errorf("URL is not found: %s", short)
}
