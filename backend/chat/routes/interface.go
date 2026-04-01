package routes

import "net/http"

type Router interface {
	Handle(pattern string, handler http.HandlerFunc)
	Serve()
}

type Handler interface{}
