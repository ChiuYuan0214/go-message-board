package routes

import (
	"general/services"
	"net/http"
)

var viewMap = MethodMapType{}

func init() {
	viewMap.put(recordView)
}

func handleView(writer http.ResponseWriter, req *http.Request) {
	setHeader(writer, "json")
	res, status := viewMap.useHandler(writer, req)
	DoResponse(res, status, writer)
}

func recordView(req *http.Request) (res interface{}, statusCode int) {
	articleId := getParam(req, "articleId")
	services.RecordView(articleId)
	return newRes("success"), http.StatusOK
}
