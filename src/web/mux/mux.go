package mux

import (
	"net/http"
)

type muxHandler struct {
	handlers    map[string]http.Handler
	handleFuncs map[string]func(resp http.ResponseWriter, req *http.Request)
}

func NewMuxHandler() *muxHandler {
	return &muxHandler{
		make(map[string]http.Handler),
		make(map[string]func(resp http.ResponseWriter, req *http.Request)),
	}
}

func (handler *muxHandler) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	urlPath := req.URL.Path
	if hl, ok := handler.handlers[urlPath]; ok {
		hl.ServeHTTP(resp, req)
		return
	}
	if fn, ok := handler.handleFuncs[urlPath]; ok {
		fn(resp, req)
		return
	}
	http.NotFound(resp, req)
}
func (hander *muxHandler) Handle(pattern string, hl http.Handler) {
	hander.handlers[pattern] = hl
}
func (handler *muxHandler) HandleFunc(pattern string, fn func(resp http.ResponseWriter, req *http.Request)) {
	handler.handleFuncs[pattern] = fn
}
