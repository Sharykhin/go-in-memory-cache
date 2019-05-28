package server

import (
	"bufio"
	"fmt"
	"net"
	"strings"

	"github.com/Sharykhin/go-in-memory-cache/server/telnet/errors"
	"github.com/Sharykhin/go-in-memory-cache/server/telnet/handler"
	"github.com/Sharykhin/go-in-memory-cache/server/telnet/logger"
	"github.com/Sharykhin/go-in-memory-cache/server/telnet/request"
)

const (
	EXIT = "exit"
)

type (
	// Server is main server struct that would serve all income messages and delegates to an
	// appropriate handler
	Server struct {
		Addr    string
		Handler handler.Handler
		Logger  logger.Logger
	}
)

// ListenAndServe starts listening income messages by tcp protocol based on a provided address
func (s Server) ListenAndServe() error {
	ln, err := net.Listen("tcp", s.Addr)
	if err != nil {
		return fmt.Errorf("could not initalize tcp connection: %v", err)
	}

	return s.Serve(ln)
}

// Serve actually serve connections and then handle them
func (s Server) Serve(ln net.Listener) error {
	defer errors.CheckDeferred(ln.Close)

	if s.Handler == nil {
		panic(errors.New("handler was not set up"))
	}

	for {
		conn, err := ln.Accept()
		if err != nil {
			return fmt.Errorf("could not accept a new connection: %v", err)
		}

		go s.handle(conn, s.Handler)
	}
}

func (s Server) handle(conn net.Conn, handler handler.Handler) {
	defer errors.CheckDeferred(conn.Close)

	buf := bufio.NewReader(conn)
	for {
		msg, err := buf.ReadString('\n')
		if err != nil {
			s.Logger.Errorf("failed to read message from a connection: %v", err)
			return
		}

		msg = strings.TrimSpace(msg)

		if msg == EXIT {
			s.Logger.Printf("connection closed.\n")
			return
		}

		r := request.NewRequestFromMessage(msg)
		handler.Serve(conn, r)
	}
}

// ListenAndServe start listening income messages and handles them
func ListenAndServe(addr string, handler handler.Handler) error {

	server := &Server{
		Addr:    addr,
		Handler: handler,
		Logger:  logger.NewTerminalLogger(),
	}

	return server.ListenAndServe()
}
