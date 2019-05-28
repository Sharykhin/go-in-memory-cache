package server

import (
	"errors"
	"fmt"
	"io"
	"net"
)

type Handler interface {
	Serve(w io.Writer, r Request)
}

type Request struct {
	Command string
	Args []string
}

type Server struct {
	Addr string
	Handler Handler
}

func (s Server) ListenAndServe() error {
	ln, err := net.Listen("tcp", s.Addr)
	if err != nil {
		return fmt.Errorf("could not initalize tcp connection: %v", err)
	}

	return s.Serve(ln)
}

func (s Server) Serve(ln net.Listener) error {
	defer ln.Close()

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

func (s Server) handle(conn net.Conn, handler Handler) {
	defer conn.Close()

	handler.Serve(conn, Request{})
}

func ListenAndServe(addr string, handler Handler) error {

	server := &Server{
		Addr:addr,
		Handler:handler,
	}

	return server.ListenAndServe()
}