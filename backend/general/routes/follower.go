package routes

import (
	"encoding/json"
	"general/service"
	"general/types"
	"net/http"

	"github.com/gin-gonic/gin"
)

type FollowerHandler struct {
	router          Router
	followerService service.Follower
}

func (fh *FollowerHandler) Run() (err error) {
	fh.router.Get("/follower", fh.get)
	fh.router.Delete("/follower", fh.remove)
	return
}

func (fh *FollowerHandler) Stop() {}

func (fh *FollowerHandler) remove(c *gin.Context) {
	var data types.FollowerData
	val, _ := c.Get("userId")
	userId := val.(uint64)
	err := json.NewDecoder(c.Request.Body).Decode(&data)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": "body format was wrong."})
		return
	}
	if userId == 0 || data.Follower == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": "userId and follower cannot be empty"})
		return
	}
	if !fh.followerService.RemoveFollower(userId, data.Follower) {
		c.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": "failed to delete data"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "success"})
}

func (fh *FollowerHandler) get(c *gin.Context) {
	userId := getUserIdFromQuery(c.Request)
	if userId == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": "userId not valid."})
		return
	}
	data := fh.followerService.GetFollowers(userId)
	c.JSON(http.StatusOK, gin.H{"status": "success", "list": data})
}
