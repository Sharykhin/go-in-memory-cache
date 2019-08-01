package cache

import (
	"fmt"
	"sync"
	"time"
)

type (
	// InMemoryStorage is a redis-like cache storage
	// that implements similar methods
	InMemoryStorage struct {
		mu     sync.RWMutex
		bucket map[string]interface{}
	}

	// Response is a general response struct with appropriate type and value
	Response struct {
		T     string
		Value interface{}
	}
)

// New is a function constructor that return a new instance of in-memory cache
func New() *InMemoryStorage {
	return &InMemoryStorage{
		mu:     sync.RWMutex{},
		bucket: make(map[string]interface{}),
	}
}

// String implements Stringer interface so this struct can be used by fmt package
func (r Response) String() string {
	return fmt.Sprintf("(%s) %v", r.T, r.Value)
}

// SET sets a new key-value. It overrides the previous value if it was set
func (s *InMemoryStorage) SET(key string, value interface{}) (*Response, error) {
	var t string
	switch value.(type) {
	case string:
		t = "string"
	case int:
		t = "integer"
	case Lists:
		t = "lists"
	case Dict:
		t = "dict"
	default:
		return nil, NewError(UnsupportedTypeCode, fmt.Sprintf("value has unsupported type %T", value))
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	s.bucket[key] = value

	return &Response{T: t, Value: value}, nil
}

// GET returns value by its key
func (s InMemoryStorage) GET(key string) interface{} {
	s.mu.RLock()
	defer s.mu.RUnlock()

	return s.bucket[key]
}

// EXPIRE set expiration time for a specific key
func (s InMemoryStorage) EXPIRE(key string, ttl int) (string, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	value, ok := s.bucket[key]
	if !ok {
		return "", NewError(KeyDoesNotExistCode, fmt.Sprintf("key %s does not exist", key))
	}

	time.AfterFunc(time.Duration(ttl), func() {
		s.mu.Lock()
		defer s.mu.Unlock()
		val := s.bucket[key]
		// TODO: ensure that value is passed properly
		if val == value {
			delete(s.bucket, key)
		}
	})

	return "OK", nil
}
