package routes

import (
	"chat/constants"
	"log"
	"net/http"
)

var _ Router = (*RouterImpl)(nil)

type RouterImpl struct{}

func (r *RouterImpl) Run() (err error) {
	return
}

func (r *RouterImpl) Stop() {}

func (r *RouterImpl) Handle(pattern string, handler http.HandlerFunc) {
	http.HandleFunc(pattern, handler)
}

func (r *RouterImpl) Serve() {
	log.Println("listening on port " + constants.PORT)
	log.Fatal(http.ListenAndServe(constants.PORT, UseCORS()))
}
