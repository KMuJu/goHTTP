package internal

import (
	"fmt"

	"github.com/kmuju/goHTTP/internal/status"
)

type Handler interface {
	ServeHTTP(ResponseWriter, *Request)
}

type HandlerFunc func(w ResponseWriter, r *Request)

func (f HandlerFunc) ServeHTTP(w ResponseWriter, r *Request) {
	f(w, r)
}

var (
	NotFoundHandler HandlerFunc = func(w ResponseWriter, r *Request) {
		fmt.Printf("Url not found: %s\n", r.URL)
		w.WriteHeader(status.NotFound)
		w.Write([]byte("Not Found"))
	}
	InternalErrorHandler HandlerFunc = func(w ResponseWriter, r *Request) {
		w.WriteHeader(status.InternalServerError)
		w.Write([]byte("Internal server Error"))
	}
)
