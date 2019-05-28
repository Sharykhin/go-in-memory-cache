package handler

import (
	"fmt"
	"io"
	"strings"

	"github.com/Sharykhin/go-in-memory-cache/server/telnet/storage"

	"github.com/Sharykhin/go-in-memory-cache/server/telnet/request"
)

type (
	// Handler is a general interface that describes method to be able to serve incoming messages properly
	Handler interface {
		Serve(w io.Writer, r *request.Request)
	}

	// CacheHandler would work around cache and use embed cache package
	CacheHandler struct {
		storage storage.Storage
	}
)

// Serve serves concrete requests
func (h CacheHandler) Serve(w io.Writer, r *request.Request) {
	fmt.Println(r)

	switch strings.ToUpper(r.Command) {
	default:
		_, _ = w.Write([]byte("command is not supported yet or invalid"))
	}
}

// NewCacheHandler is a constructor that returns cache handler
func NewCacheHandler() *CacheHandler {
	return &CacheHandler{
		storage: storage.NewMemoryCache(),
	}
}
