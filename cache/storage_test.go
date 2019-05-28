package cache

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestInMemoryStorage_SET(t *testing.T) {
	cache := New()

	res, err := cache.SET("test", 10)
	if err != nil {
		t.Errorf("failed to set key: %v", err)
	}

	assert.Equal(t, "integer", res.T)
	assert.Equal(t, 10, res.Value)

}

func TestInMemoryStorage_GET(t *testing.T) {
	cache := New()

	val := cache.GET("nonexisting")
	assert.Equal(t, val, nil)

	_, _ = cache.SET("testkey", "hello world")

	val = cache.GET("testkey")
	assert.Equal(t, "hello world", val)
}

func TestInMemoryStorage_EXPIRE(t *testing.T) {
	cache := New()

	_, _ = cache.SET("testkey", "hello world")

	res, err := cache.EXPIRE("testkey", 1)
	if err != nil {
		t.Errorf("faield to set TTL: %v", err)
	}

	assert.Equal(t, "OK", res)

	time.Sleep(1001 * time.Millisecond)
	value := cache.GET("testkey")
	assert.Equal(t, nil, value)

	_, err = cache.EXPIRE("nonexisting", 1)

	assert.NotNil(t, err)
	cacheErr, _ := err.(*Error)
	assert.Equal(t, KeyDoesNotExistCode, cacheErr.Code)
}
