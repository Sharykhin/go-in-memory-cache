package cache

import (
	"fmt"
	"strings"
	"sync"
	"time"
)

const (
	SuccessResponse = "OK"
	NoneResponse    = "none"
)

type (
	InMemoryStorage struct {
		mu     sync.RWMutex
		bucket map[string]interface{}
	}

	// Response
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

// SET
func (s *InMemoryStorage) SET(key string, value interface{}) (string, error) {
	switch value.(type) {
	case string:
		fmt.Println("string", value)

	case int:
		fmt.Println("integer", value)

	case Lists:
		fmt.Println("lists", value)

	case Dict:
		fmt.Println("dict", value)

	default:
		return "", NewError(UnsupportedTypeCode, fmt.Sprintf("value has unsupported type %T", value))
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	s.bucket[key] = value

	return SuccessResponse, nil
}

// GET returns value by its key
// If key does not exist it will return nil
func (s *InMemoryStorage) GET(key string) interface{} {
	s.mu.RLock()
	defer s.mu.RUnlock()

	return s.bucket[key]
}

func (s *InMemoryStorage) EXPIRE(key string, ttl int) (string, error) {
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

	return SuccessResponse, nil
}

func (s *InMemoryStorage) TYPE(key string) string {
	s.mu.RLock()
	defer s.mu.RUnlock()

	value, ok := s.bucket[key]
	if !ok {
		return NoneResponse
	}

	// TODO: move all type representations into constants
	switch value.(type) {
	case string:
		return "string"
	case int:
		return "integer"
	case Lists:
		return "lists"
	case Dict:
		return "dict"
	default:
		return NoneResponse
	}
}

func (s *InMemoryStorage) LPUSH(key string, value string) (int, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	list, ok := s.bucket[key]
	if !ok {
		s.bucket[key] = Lists{value}
		return 1, nil
	}

	l, ok := list.(Lists)
	if !ok {
		return 0, NewError(CorruptedListCode, "could not convert list into internal representation")
	}

	l = append([]string{value}, l...)
	s.bucket[key] = l

	return len(l), nil
}

func (s *InMemoryStorage) RPUSH(key string, value string) (int, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	list, ok := s.bucket[key]
	if !ok {
		s.bucket[key] = Lists{value}
		return 1, nil
	}

	l, ok := list.(Lists)
	if !ok {
		return 0, NewError(CorruptedListCode, "could not convert list into internal representation")
	}

	l = append(l, value)
	s.bucket[key] = l

	return len(l), nil
}

func (s *InMemoryStorage) LRANGE(key string, start, end int) (res []string, err error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	defer func() {
		if r := recover(); r != nil {
			if strings.Contains(fmt.Sprintf("%s", r), "out of range") {
				err = NewError(SliceBoundsOutOfRange, "index our of range")
			} else {
				panic(r)
			}
		}
	}()

	list, ok := s.bucket[key]
	if !ok {
		return nil, NewError(EmptyList, "list is empty")
	}

	l, ok := list.(Lists)
	if !ok {
		return nil, NewError(CorruptedListCode, "could not convert list into internal representation")
	}
	if end < 0 {
		end = len(l) + 1 + end
	}

	return l[start:end], nil
}

func (s *InMemoryStorage) RPOP(key string) (string, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	list, ok := s.bucket[key]
	if !ok {
		return "", NewError(EmptyList, "list is empty")
	}

	l, ok := list.(Lists)
	if !ok {
		return "", NewError(CorruptedListCode, "could not convert list into internal representation")
	}

	item := l[len(l)-1]
	s.bucket[key] = l[:len(l)-1]

	return item, nil
}
