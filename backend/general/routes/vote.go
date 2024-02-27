package routes

import (
	"general/services"
	"general/utils"
	"net/http"
)

type NewVoteData struct {
	SourceId int64  `json:"sourceId"` // <articleId> | <commentId>
	Score    int16  `json:"score"`
	VoteType string `json:"voteType"` // "article" | "comment"
}

type UpdateVoteData struct {
	VoteId int64 `json:"voteId"`
	Score  int16 `json:"score"`
}

var voteMap = MethodMapType{}

func init() {
	voteMap.post(newVote).put(updateVote)
}

func handleVote(writer http.ResponseWriter, req *http.Request) {
	setHeader(writer, "json")
	res, status := voteMap.useHandler(writer, req)
	DoResponse(res, status, writer)
}

func newVote(req *http.Request) (res interface{}, statusCode int) {
	userId := getUserIdFromContext(req)
	data := &NewVoteData{}
	message, status := utils.ParseBody(req.Body, data)
	if message != "" {
		return newRes("fail").message(message), status
	}

	if data.SourceId == 0 || !utils.ContainsString([]string{"article", "comment"}, data.VoteType) {
		return newRes("fail").message("sourceId cannot be empty, voteType should be either article or comment."), http.StatusBadRequest
	}
	if data.Score < -1 || data.Score > 1 {
		return newRes("fail").message("source can only be 1 or 0 or -1."), http.StatusBadRequest
	}

	message, voteIdOrStatus := services.Vote(userId, data.SourceId, data.Score, &data.VoteType)
	if message != "" {
		return newRes("fail").message(message), int(voteIdOrStatus)
	}

	return newRes("success").setId(voteIdOrStatus), http.StatusOK
}

func updateVote(req *http.Request) (res interface{}, statusCode int) {
	userId := getUserIdFromContext(req)
	data := &UpdateVoteData{}
	message, status := utils.ParseBody(req.Body, data)
	if message != "" {
		return newRes("fail").message(message), status
	}

	if data.VoteId == 0 {
		return newRes("fail").message("voteId cannot be empty"), http.StatusBadRequest
	}
	if data.Score < -1 || data.Score > 1 {
		return newRes("fail").message("source can only be 1 or 0 or -1."), http.StatusBadRequest
	}

	if !services.UpdateVote(userId, data.VoteId, data.Score) {
		return newRes("fail").message("userId or voteId incorrect."), http.StatusBadRequest
	}

	return newRes("success"), http.StatusOK
}
