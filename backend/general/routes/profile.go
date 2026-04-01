package routes

import (
	"general/routes/middleware"
	"general/service"
	"general/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ProfileHandler struct {
	router         Router
	profileService service.Profile
}

func (ph *ProfileHandler) Run() (err error) {
	ph.router.Get("/profile", ph.get)
	return
}

func (ph *ProfileHandler) Stop() {}

func (ph *ProfileHandler) get(c *gin.Context) {
	userId := middleware.GetUserIdFromHeader(c)

	if userId != 0 {
		profile, status := ph.profileService.GetProfileWithId(userId)
		if profile != nil {
			c.JSON(status, gin.H{"status": "success", "data": *profile})
			return
		}
		if status == http.StatusBadRequest {
			c.JSON(status, gin.H{"status": "fail", "message": "invalid input."})
			return
		}
		c.JSON(status, gin.H{"status": "fail", "message": "something went wrong."})
		return
	}

	id := utils.IsAuth(c.Request)
	profile, status := ph.profileService.GetProfileWithToken(id)
	if profile != nil {
		c.JSON(status, gin.H{"status": "success", "data": *profile})
		return
	}
	c.JSON(status, gin.H{"status": "fail", "message": "something went wrong."})
}
