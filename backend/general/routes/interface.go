package routes

import "github.com/gin-gonic/gin"

type Router interface {
	Get(path string, handlers ...gin.HandlerFunc)
	Post(path string, handlers ...gin.HandlerFunc)
	Put(path string, handlers ...gin.HandlerFunc)
	Delete(path string, handlers ...gin.HandlerFunc)
	Serve()
}

type Handler interface {
}
