package handler

import (
	"fmt"
	"io"

	"github.com/Sharykhin/go-in-memory-cache/cache"

	"github.com/Sharykhin/go-in-memory-cache/server/telnet/request"
)

type (
	// Handler is a general interface that describes method to be able to serve incoming messages properly
	Handler interface {
		Serve(w io.Writer, r *request.Request)
	}

	Storage interface {
		SET(key string, value interface{}) (string, error)
		GET(key string) interface{}
	}

	// CacheHandler would work around cache and use embed cache package
	CacheHandler struct {
		storage Storage
	}
)

// Serve serves concrete requests
func (h CacheHandler) Serve(w io.Writer, r *request.Request) {
	fmt.Println(r)
}

// NewCacheHandler is a constructor that returns cache handler
func NewCacheHandler() *CacheHandler {
	return &CacheHandler{
		storage: cache.New(),
	}
}
