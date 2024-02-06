package routes

import (
	"general/services"
	"general/types"
	"net/http"
)

var articlesMap = MethodMapType{}

func init() {
	articlesMap.get(getArticles)
}

func handleArticles(writer http.ResponseWriter, req *http.Request) {
	setContentType(writer, "json")
	res, status := articlesMap.useHandler(req)
	DoResponse(res, status, writer)
}

func getArticles(req *http.Request) (res interface{}, statusCode int) {
	page, size := getPageSize(req)
	listType := getParam(req, "type")
	tag := getParam(req, "tag")
	userId := getUserIdFromQuery(req)
	var articles *[]types.ArticleListData
	switch listType {
	case "view":
		articles = services.GetViewList(page, size)
	case "hot":
		articles = services.GetHotList(page, size)
	case "profile":
		articles = services.GetProfileList(page, size, userId)
	case "tag":
		articles = services.GetTagList(page, size, tag)
	default:
		articles = services.GetNewestList(page, size)
	}

	return newRes("success").setList("list", *articles), http.StatusOK
}
