package routes

import (
	"general/routes/middleware"
	"general/service"
	"general/types"
	"general/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CommentHandler struct {
	router         Router
	commentService service.Comment
}

func (ch *CommentHandler) Run() (err error) {
	ch.router.Post("/comment", middleware.Auth(), ch.add)
	ch.router.Put("/comment", middleware.Auth(), ch.update)
	ch.router.Delete("/comment", middleware.Auth(), ch.delete)
	return
}

func (ch *CommentHandler) Stop() {}

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

	commentId := ch.commentService.AddComment(userId, data)
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

	message, status = ch.commentService.UpdateComment(userId, data)
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
	message, status := ch.commentService.DeleteComment(userId, commentId)
	if message != "" {
		c.JSON(status, gin.H{"status": "fail", "message": message})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "success"})
}
