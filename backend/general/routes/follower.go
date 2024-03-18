package routes

import (
	"encoding/json"
	"general/services"
	"general/types"
	"net/http"

	"github.com/gin-gonic/gin"
)

func initFollower(router *gin.Engine) {
	fh := FollowerHandler{}
	router.GET("/follower", fh.get)
	router.DELETE("/follower", fh.remove)
}

type FollowerHandler struct{}

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
	if !services.RemoveFollower(userId, data.Follower) {
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
	data := services.GetFollowers(userId)
	c.JSON(http.StatusOK, gin.H{"status": "success", "list": data})
}
