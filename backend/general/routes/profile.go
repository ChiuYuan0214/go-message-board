package routes

import (
	"general/routes/middleware"
	"general/services"
	"general/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func initProfile(router *gin.Engine) {
	ph := ProfileHandler{}
	router.GET("/profile", ph.get)
}

type ProfileHandler struct{}

func (ph *ProfileHandler) get(c *gin.Context) {
	userId := middleware.GetUserIdFromHeader(c)

	if userId != 0 {
		profile, status := services.GetProfileWithId(userId)
		if profile != nil {
			c.JSON(status, gin.H{"status": "success", "data": *profile})
			return
		}
		if status == http.StatusBadRequest {
			c.JSON(status, gin.H{"status": "fail", "message": "invalid input."})
			return
		}
		c.JSON(status, gin.H{"status": "fail", "message": "something went wrong."})
	}

	id := utils.IsAuth(c.Request)
	profile, status := services.GetProfileWithToken(id)
	if profile != nil {
		c.JSON(status, gin.H{"status": "success", "data": *profile})
		return
	}
	c.JSON(status, gin.H{"status": "fail", "message": "something went wrong."})
}
