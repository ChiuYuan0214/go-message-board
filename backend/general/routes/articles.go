package routes

import (
	"general/services"
	"general/types"
	"general/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func initArticles(router *gin.Engine) {
	ah := ArticlesHandler{}
	router.GET("/articles", ah.get)
}

type ArticlesHandler struct{}

func (ah *ArticlesHandler) get(c *gin.Context) {
	page, size := getPageSize(c.Request)
	listType := getParam(c.Request, "type")
	tag := getParam(c.Request, "tag")
	userId := getUserIdFromQuery(c.Request)
	selfUserId := utils.IsAuth(c.Request)
	var articles []types.ArticleListData
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

	c.JSON(http.StatusOK, gin.H{"status": "success", "list": articles})
}
