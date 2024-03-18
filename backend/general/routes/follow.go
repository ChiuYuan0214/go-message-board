package routes

import (
	"encoding/json"
	"general/routes/middleware"
	"general/services"
	"general/types"
	"net/http"

	"github.com/gin-gonic/gin"
)

func initFollow(router *gin.Engine) {
	fh := FollowHandler{}
	router.POST("/follow", middleware.Auth(), fh.add)
	router.DELETE("/follow", middleware.Auth(), fh.remove)
}

type FollowHandler struct{}

func (fh *FollowHandler) add(c *gin.Context) {
	var data types.FollowData
	err := json.NewDecoder(c.Request.Body).Decode(&data)
	val, _ := c.Get("userId")
	userId := val.(uint64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": "body format was wrong."})
		return
	}
	if userId == 0 || data.Followee == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": "userId and followee cannot be empty"})
		return
	}
	if !services.AddFollow(userId, data.Followee) {
		c.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": "failed to insert data"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "success"})
}

func (fh *FollowHandler) remove(c *gin.Context) {
	var data types.FollowData
	err := json.NewDecoder(c.Request.Body).Decode(&data)
	val, _ := c.Get("userId")
	userId := val.(uint64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": "body format was wrong."})
		return
	}
	if userId == 0 || data.Followee == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": "userId and followee cannot be empty"})
		return
	}
	if !services.RemoveFollow(userId, data.Followee) {
		c.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": "failed to delete data"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "success"})
}
