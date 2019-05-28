package cache

import (
	"fmt"
	"sync"
	"time"
)

type (
	InMemoryStorage struct {
		mu     sync.RWMutex
		bucket map[string]interface{}
	}

	Response struct {
		T     string
		Value interface{}
	}
)

func New() *InMemoryStorage {
	return &InMemoryStorage{
		mu:     sync.RWMutex{},
		bucket: make(map[string]interface{}),
	}
}

func (r Response) String() string {
	return fmt.Sprintf("(%s) %v", r.T, r.Value)
}

func (s *InMemoryStorage) SET(key string, value interface{}) (*Response, error) {
	var t string
	switch value.(type) {
	case string:
		fmt.Println("string", value)
		t = "string"
	case int:
		fmt.Println("integer", value)
		t = "integer"
	case Lists:
		fmt.Println("lists", value)
		t = "lists"
	case Dict:
		fmt.Println("dict", value)
		t = "dict"
	default:
		return nil, NewError(UnsupportedTypeCode, fmt.Sprintf("value has unsupported type %T", value))
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	s.bucket[key] = value

	return &Response{T: t, Value: value}, nil
}

func (s InMemoryStorage) GET(key string) interface{} {
	s.mu.RLock()
	defer s.mu.RUnlock()

	return s.bucket[key]
}

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
