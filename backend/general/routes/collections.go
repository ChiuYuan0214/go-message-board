package routes

import (
	"encoding/json"
	"general/routes/middleware"
	"general/services"
	"general/types"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func initCollections(router *gin.Engine) {
	ch := CollectionsHandler{}
	router.GET("/collections", middleware.Auth(), ch.get)
	router.POST("/collections", middleware.Auth(), ch.add)
	router.DELETE("/collections", middleware.Auth(), ch.delete)
}

type CollectionsHandler struct{}

func (ch *CollectionsHandler) get(c *gin.Context) {
	val, _ := c.Get("userId")
	userId := val.(uint64)
	page, size := getPageSize(c.Request)
	if userId == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": "userId not valid."})
		return
	}
	data := services.GetCollections(userId, page, size)
	c.JSON(http.StatusOK, gin.H{"status": "success", "list": data})
}

func (ch *CollectionsHandler) add(c *gin.Context) {
	val, _ := c.Get("userId")
	userId := val.(uint64)
	var body types.WriteCollectionData
	err := json.NewDecoder(c.Request.Body).Decode(&body)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": "body format was wrong."})
		return
	}
	if userId == 0 || body.ArticleId == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": "userId and articleId cannot be empty."})
		return
	}
	if !services.AddCollection(userId, body.ArticleId) {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "fail", "message": "failed to add collection."})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "success"})
}

func (ch *CollectionsHandler) delete(c *gin.Context) {
	val, _ := c.Get("userId")
	userId := val.(uint64)
	var body types.WriteCollectionData
	err := json.NewDecoder(c.Request.Body).Decode(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": "body format was wrong."})
		return
	}
	if userId == 0 || body.ArticleId == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": "userId and articleId cannot be empty."})
		return
	}
	if !services.RemoveCollection(userId, body.ArticleId) {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "fail", "message": "failed to remove collection."})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "success"})
}
