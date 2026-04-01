package routes

import (
	"general/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CommentsHandler struct {
	router          Router
	commentsService service.Comments
}

func (ch *CommentsHandler) Run() (err error) {
	ch.router.Get("/comments", ch.get)
	return
}

func (ch *CommentsHandler) Stop() {}

func (ch *CommentsHandler) get(c *gin.Context) {
	id := getParam(c.Request, "articleId")
	articleId, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": "articleId format incorrect."})
		return
	}
	data := ch.commentsService.GetComments(articleId)
	c.JSON(http.StatusOK, gin.H{"status": "success", "list": data})
}
