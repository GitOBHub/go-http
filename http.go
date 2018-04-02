package http

import (
	"log"
)

type Handler interface {
	ServeHTTP(w ResponseWriter, r *Request)
}

type HandlerFunc func(ResponseWriter, *Request)

type ServeMux struct {
	handlers map[string]HandlerFunc
}

var defaultServeMux *ServeMux

const (
	StatusOK       int = 200
	StatusNotFound int = 404
)

func ListenAndServe(addr string, h Handler) error {
	srv := Server{addr, h}
	return srv.ListenAndServe()
}

func NewServeMux() *ServeMux {
	mux := new(ServeMux)
	mux.handlers = make(map[string]HandlerFunc)
	return mux
}

func HandleFunc(pattern string, handler func(ResponseWriter, *Request)) {
	defaultServeMux.handlers[pattern] = HandlerFunc(handler)
}

func (f HandlerFunc) ServeHTTP(w ResponseWriter, r *Request) {
	f(w, r)
}

func (m ServeMux) ServeHTTP(w ResponseWriter, r *Request) {
	handler, ok := m.handlers[r.URL.Path]
	if !ok {
		NotFound(w, r)
		return
	}
	log.Print("ServeMux.ServeHTTP: Found")
	handler.ServeHTTP(w, r)
}

func init() {
	defaultServeMux = NewServeMux()
}

func NotFound(w ResponseWriter, r *Request) {
	w.WriteHeader(StatusNotFound)
}
