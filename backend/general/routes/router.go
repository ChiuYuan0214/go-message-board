package routes

import (
	"fmt"
	"general/constants"
	"general/routes/middleware"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

var _ Router = (*RouterImpl)(nil)

type RouterImpl struct {
	router *gin.Engine
}

func (r *RouterImpl) Run() (err error) {
	r.router = gin.Default()
	r.router.Use(middleware.Cors)
	return
}

func (r *RouterImpl) Stop() {}

func (r *RouterImpl) Get(path string, handlers ...gin.HandlerFunc) {
	r.router.GET(path, handlers...)
}

func (r *RouterImpl) Post(path string, handlers ...gin.HandlerFunc) {
	r.router.POST(path, handlers...)
}

func (r *RouterImpl) Put(path string, handlers ...gin.HandlerFunc) {
	r.router.PUT(path, handlers...)
}

func (r *RouterImpl) Delete(path string, handlers ...gin.HandlerFunc) {
	r.router.DELETE(path, handlers...)
}

func (r *RouterImpl) Serve() {
	r.router.Run()
	fmt.Printf("Server listening on :%s...\n", constants.PORT)
	log.Println(http.ListenAndServe(fmt.Sprintf(":%s", constants.PORT), nil))
}
