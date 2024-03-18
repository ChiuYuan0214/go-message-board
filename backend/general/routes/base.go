package routes

import (
	"general/routes/middleware"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var db *gorm.DB

func UsePool(DB *gorm.DB) {
	db = DB
}

func InitRouter(router *gin.Engine) {
	router.Use(middleware.Cors)
	initArticle(router)
	initArticles(router)
	initComment(router)
	initComments(router)
	initProfile(router)
	initCollections(router)
	initVote(router)
	initView(router)
	initFollow(router)
	initFollower(router)
}
