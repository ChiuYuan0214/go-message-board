package routes

import (
	"general/services"
	"net/http"
	"strconv"
)

var commentsMap MethodMapType = map[string]HandlerType{}

func init() {
	commentsMap.get(getComments)
}

func handleComments(writer http.ResponseWriter, req *http.Request) {
	setHeader(writer, "json")
	res, status := commentsMap.useHandler(writer, req)
	DoResponse(res, status, writer)
}

func getComments(req *http.Request) (res interface{}, statusCode int) {
	id := getParam(req, "articleId")
	articleId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return newRes("fail").message("articleId format incorrect."), http.StatusBadRequest
	}
	data := services.GetComments(articleId)
	return newRes("success").setList("list", *data), http.StatusOK
}
