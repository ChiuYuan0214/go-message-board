package routes

import (
	"general/services"
	"general/types"
	"general/utils"
	"net/http"
)

var articlesMap = MethodMapType{}

func init() {
	articlesMap.get(getArticles)
}

func handleArticles(writer http.ResponseWriter, req *http.Request) {
	setHeader(writer, "json")
	res, status := articlesMap.useHandler(writer, req)
	DoResponse(res, status, writer)
}

func getArticles(req *http.Request) (res interface{}, statusCode int) {
	page, size := getPageSize(req)
	listType := getParam(req, "type")
	tag := getParam(req, "tag")
	userId := getUserIdFromQuery(req)
	selfUserId := utils.IsAuth(req)
	var articles *[]types.ArticleListData
	switch listType {
	case "view":
		articles = services.GetViewList(page, size, userId)
	case "hot":
		articles = services.GetHotList(page, size, userId)
	case "profile":
		articles = services.GetProfileList(page, size, userId, selfUserId)
	case "tag":
		articles = services.GetTagList(page, size, tag)
	default:
		articles = services.GetNewestList(page, size, userId)
	}

	return newRes("success").setList("list", *articles), http.StatusOK
}
