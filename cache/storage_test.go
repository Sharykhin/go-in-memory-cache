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

	assert.Equal(t, SuccessResponse, res)
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

func TestInMemoryStorage_TYPE(t *testing.T) {
	cache := New()

	_, _ = cache.SET("testkeystring", "hello world")
	_, _ = cache.SET("testkeyint", 10)

	res := cache.TYPE("testkeystring")

	assert.Equal(t, "string", res)

	res = cache.TYPE("testkeyint")

	assert.Equal(t, "integer", res)

	res = cache.TYPE("nonexisting")
	assert.Equal(t, NoneResponse, res)
}

func TestInMemoryStorage_LPUSH(t *testing.T) {
	cache := New()

	num, err := cache.LPUSH("testlist", "hello")
	if err != nil {
		t.Errorf("failed to push item into a list: %v", err)
	}

	assert.Equal(t, 1, num)

	num, err = cache.LPUSH("testlist", "world", "!!!")
	if err != nil {
		t.Errorf("failed to push item into a list: %v", err)
	}

	assert.Equal(t, 3, num)
}

func TestInMemoryStorage_RPUSH(t *testing.T) {
	cache := New()

	num, err := cache.RPUSH("testlist", "hello")
	if err != nil {
		t.Errorf("failed to push item into a list: %v", err)
	}

	assert.Equal(t, 1, num)

	num, err = cache.RPUSH("testlist", "world")
	if err != nil {
		t.Errorf("failed to push item into a list: %v", err)
	}

	assert.Equal(t, 2, num)
}

func TestInMemoryStorage_LRANGE(t *testing.T) {
	cache := New()
	_, _ = cache.RPUSH("testlist", "hello")
	_, _ = cache.RPUSH("testlist", "world")

	_, err := cache.LRANGE("testlist", -2, 2)
	cacheErr := err.(*Error)
	assert.Equal(t, SliceBoundsOutOfRange, cacheErr.Code)

	res, err := cache.LRANGE("testlist", 0, 2)
	assert.Equal(t, []string{"hello", "world"}, res)

	res, err = cache.LRANGE("testlist", 0, -1)
	assert.Equal(t, []string{"hello", "world"}, res)

	res, err = cache.LRANGE("testlist", 0, -2)
	assert.Equal(t, []string{"hello"}, res)
}

func TestInMemoryStorage_RPOP(t *testing.T) {
	cache := New()
	_, _ = cache.RPUSH("testlist", "hello")
	_, _ = cache.RPUSH("testlist", "world")

	res, _ := cache.LRANGE("testlist", 0, -1)
	assert.Equal(t, []string{"hello", "world"}, res)

	item, err := cache.RPOP("testlist")
	if err != nil {
		t.Errorf("failed to pop the last element: %v", err)
	}

	assert.Equal(t, "world", item)
	res, _ = cache.LRANGE("testlist", 0, -1)
	assert.Equal(t, []string{"hello"}, res)
}
