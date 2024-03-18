package routes

import (
	"general/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func getFollows(c *gin.Context) {
	userId := getUserIdFromQuery(c.Request)
	if userId == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": "userId not valid."})
		return
	}
	data := services.GetFollows(userId)
	c.JSON(http.StatusOK, gin.H{"status": "success", "list": data})
}
