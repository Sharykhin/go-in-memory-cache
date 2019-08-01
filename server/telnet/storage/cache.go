package storage

import (
	"fmt"

	"github.com/Sharykhin/go-in-memory-cache/cache"
)

type (
	Storage interface {
		SetInt(key string, value int) error
		GetInt(key string) (int, error)
	}

	MemoryCache struct {
		storage *cache.InMemoryStorage
	}
)

func NewMemoryCache() *MemoryCache {
	return &MemoryCache{
		storage: cache.New(),
	}
}

func (c MemoryCache) SetInt(key string, value int) error {
	_, err := c.storage.SET(key, value)
	if err != nil {
		return fmt.Errorf("could not set value: %v", err)
	}

	return nil
}

func (c MemoryCache) GetInt(key string) (int, error) {
	res := c.storage.GET(key)
	if val, ok := res.(int); !ok {
		return 0, fmt.Errorf("failed to get integer, value has %T type", val)
	} else {
		return val, nil
	}
}
