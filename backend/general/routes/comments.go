package routes

import (
	"general/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func initComments(router *gin.Engine) {
	ch := CommentsHandler{}
	router.GET("/comments", ch.get)
}

type CommentsHandler struct{}

func (ch *CommentsHandler) get(c *gin.Context) {
	id := getParam(c.Request, "articleId")
	articleId, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": "articleId format incorrect."})
		return
	}
	data := services.GetComments(articleId)
	c.JSON(http.StatusOK, gin.H{"status": "success", "list": data})
}
