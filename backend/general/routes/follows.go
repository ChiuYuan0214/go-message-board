package routes

import (
	"general/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type FollowsHandler struct {
	router        Router
	followsService service.Follows
}

func (fh *FollowsHandler) Run() (err error) {
	fh.router.Get("/follows", fh.get)
	return
}

func (fh *FollowsHandler) Stop() {}

func (fh *FollowsHandler) get(c *gin.Context) {
	userId := getUserIdFromQuery(c.Request)
	if userId == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": "userId not valid."})
		return
	}
	data := fh.followsService.GetFollows(userId)
	c.JSON(http.StatusOK, gin.H{"status": "success", "list": data})
}
