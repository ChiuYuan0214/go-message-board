package routes

import (
	"general/services"
	"net/http"
)

var followsMap = MethodMapType{}

func init() {
	followsMap.get(getFollows)
}

func handleFollows(writer http.ResponseWriter, req *http.Request) {
	setHeader(writer, "json")
	res, status := followsMap.useHandler(writer, req)
	DoResponse(res, status, writer)
}

func getFollows(req *http.Request) (res interface{}, statusCode int) {
	userId := getUserIdFromQuery(req)
	if userId == 0 {
		return newRes("fail").message("userId not valid."), http.StatusBadRequest
	}
	data := services.GetFollows(userId)
	return newRes("success").setList("list", *data), http.StatusOK
}
