package handler

import (
	"fmt"
	"io"

	"github.com/Sharykhin/go-in-memory-cache/server/telnet/request"
)

type (
	// Handler is a general interface that describes method to be able to serve incoming messages properly
	Handler interface {
		Serve(w io.Writer, r *request.Request)
	}

	// CacheHandler would work around cache and use embed cache package
	CacheHandler struct {
	}
)

// Serve serves concrete requests
func (h CacheHandler) Serve(w io.Writer, r *request.Request) {
	fmt.Println(r)
}

// NewCacheHandler is a constructor that returns cache handler
func NewCacheHandler() *CacheHandler {
	return &CacheHandler{}
}
