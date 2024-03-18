package routes

import (
	"general/routes/middleware"
	"general/services"
	"general/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type NewVoteData struct {
	SourceId uint64 `json:"sourceId"` // <articleId> | <commentId>
	Score    int8   `json:"score"`
	VoteType string `json:"voteType"` // "article" | "comment"
}

type UpdateVoteData struct {
	VoteId uint64 `json:"voteId"`
	Score  int8   `json:"score"`
}

func initVote(router *gin.Engine) {
	vh := VoteHandler{}
	router.POST("/vote", middleware.Auth(), vh.add)
	router.PUT("/vote", middleware.Auth(), vh.update)
}

type VoteHandler struct{}

func (vh *VoteHandler) add(c *gin.Context) {
	val, _ := c.Get("userId")
	userId := val.(uint64)
	data := &NewVoteData{}
	message, status := utils.ParseBody(c.Request.Body, data)
	if message != "" {
		c.JSON(status, gin.H{"status": "fail", "message": message})
		return
	}
	if data.SourceId == 0 || !utils.ContainsString([]string{"article", "comment"}, data.VoteType) {
		c.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": "sourceId cannot be empty, voteType should be either article or comment."})
		return
	}
	if data.Score < -1 || data.Score > 1 {
		c.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": "source can only be 1 or 0 or -1."})
		return
	}
	message, voteIdOrStatus := services.Vote(userId, data.SourceId, data.Score, &data.VoteType)
	if message != "" {
		c.JSON(int(voteIdOrStatus), gin.H{"status": "fail", "message": message})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "success", "id": voteIdOrStatus})
}

func (vh *VoteHandler) update(c *gin.Context) {
	val, _ := c.Get("userId")
	userId := val.(uint64)
	data := &UpdateVoteData{}
	message, status := utils.ParseBody(c.Request.Body, data)
	if message != "" {
		c.JSON(status, gin.H{"status": "fail", "message": message})
		return
	}
	if data.VoteId == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": "voteId cannot be empty"})
		return
	}
	if data.Score < -1 || data.Score > 1 {
		c.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": "source can only be 1 or 0 or -1."})
		return
	}
	if !services.UpdateVote(userId, data.VoteId, data.Score) {
		c.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": "userId or voteId incorrect."})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "success"})
}
