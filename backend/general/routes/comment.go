package routes

import (
	"general/routes/middleware"
	"general/services"
	"general/types"
	"general/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func initComment(router *gin.Engine) {
	ch := CommentHandler{}
	router.POST("/comment", middleware.Auth(), ch.add)
	router.PUT("/comment", middleware.Auth(), ch.update)
	router.DELETE("/comment", middleware.Auth(), ch.delete)
}

type CommentHandler struct{}

func (ch *CommentHandler) add(c *gin.Context) {
	val, _ := c.Get("userId")
	userId := val.(uint64)
	data := &types.AddCommentData{}
	message, status := utils.ParseBody(c.Request.Body, data)
	if message != "" {
		c.JSON(status, gin.H{"status": "fail", "message": message})
		return
	}
	if data.ArticleId == 0 || data.Title == "" || data.Content == "" {
		c.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": "articleId, title and content cannot be empty"})
		return
	}

	commentId := services.AddComment(userId, data)
	if commentId == 0 {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "fail", "message": "something went wrong."})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "success", "id": commentId})
}

func (ch *CommentHandler) update(c *gin.Context) {
	val, _ := c.Get("userId")
	userId := val.(uint64)
	data := &types.UpdateCommentData{}
	message, status := utils.ParseBody(c.Request.Body, data)
	if message != "" {
		c.JSON(status, gin.H{"status": "fail", "message": message})
		return
	}

	message, status = services.UpdateComment(userId, data)
	if message != "" {
		c.JSON(status, gin.H{"status": "fail", "message": message})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "success"})
}

func (ch *CommentHandler) delete(c *gin.Context) {
	val, _ := c.Get("userId")
	userId := val.(uint64)
	commentId := getParam(c.Request, "commentId")
	message, status := services.DeleteComment(userId, commentId)
	if message != "" {
		c.JSON(status, gin.H{"status": "fail", "message": message})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "success"})
}
