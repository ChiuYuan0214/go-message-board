package routes

import (
	"fmt"
	"log"
	"net/http"
	"security/constants"
)

var _ Router = (*RouterImpl)(nil)

type RouterImpl struct{}

func (r *RouterImpl) Handle(pattern string, handler http.HandlerFunc) {
	http.HandleFunc(pattern, handler)
}

func (r *RouterImpl) HandleStatic() {
	fs := http.FileServer(http.Dir("./uploads/images/"))
	http.Handle("/uploads/images/", http.StripPrefix("/uploads/images/", fs))
}

func (r *RouterImpl) Serve() {
	fmt.Printf("Server listening on :%s...\n", constants.PORT)
	log.Println(http.ListenAndServe(fmt.Sprintf(":%s", constants.PORT), nil))
}
