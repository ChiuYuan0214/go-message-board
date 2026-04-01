package routes

import (
	"general/service"
	"general/types"
	"general/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ArticlesHandler struct {
	router          Router
	articlesService service.Articles
}

func (ah *ArticlesHandler) Run() (err error) {
	ah.router.Get("/articles", ah.get)
	return
}

func (ah *ArticlesHandler) Stop() {}

func (ah *ArticlesHandler) get(c *gin.Context) {
	page, size := getPageSize(c.Request)
	listType := getParam(c.Request, "type")
	tag := getParam(c.Request, "tag")
	userId := getUserIdFromQuery(c.Request)
	selfUserId := utils.IsAuth(c.Request)
	var articles []types.ArticleListData
	switch listType {
	case "view":
		articles = ah.articlesService.GetViewList(page, size, userId)
	case "hot":
		articles = ah.articlesService.GetHotList(page, size, userId)
	case "profile":
		articles = ah.articlesService.GetProfileList(page, size, userId, selfUserId)
	case "tag":
		articles = ah.articlesService.GetTagList(page, size, tag)
	default:
		articles = ah.articlesService.GetNewestList(page, size, userId)
	}

	c.JSON(http.StatusOK, gin.H{"status": "success", "list": articles})
}
