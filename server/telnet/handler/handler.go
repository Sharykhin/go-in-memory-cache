package handler

import "io"

type (
	Handler interface {
		Serve(w io.Writer, r *request.Request)
	}

	MainHandler struct {

	}
)

func (h MainHandler) Serve(w io.Writer, r *request.Request) {

}

func NewMainHandler() *MainHandler {
	return &MainHandler{}
}