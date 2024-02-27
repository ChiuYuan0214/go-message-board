package routes

import (
	"general/services"
	"general/types"
	"general/utils"
	"net/http"
)

var commentMap = MethodMapType{}

func init() {
	commentMap.put(authMethod(updateComment)).post(authMethod(addComment)).delete(authMethod(deleteComment))
}

func handleComment(writer http.ResponseWriter, req *http.Request) {
	setHeader(writer, "json")
	res, status := commentMap.useHandler(writer, req)
	DoResponse(res, status, writer)
}

func addComment(req *http.Request) (res interface{}, statusCode int) {
	userId := getUserIdFromContext(req)
	data := &types.AddCommentData{}
	message, status := utils.ParseBody(req.Body, data)
	if message != "" {
		return newRes("fail").message(message), status
	}
	if data.ArticleId == 0 || data.Title == "" || data.Content == "" {
		return newRes("fail").message("articleId, title and content cannot be empty"), http.StatusBadRequest
	}

	commentId := services.AddComment(userId, data)
	if commentId == 0 {
		return newRes("fail").message("something went wrong."), http.StatusInternalServerError
	}

	return newRes("success").setId(commentId), http.StatusOK
}

func updateComment(req *http.Request) (res interface{}, statusCode int) {
	userId := getUserIdFromContext(req)
	data := &types.UpdateCommentData{}
	message, status := utils.ParseBody(req.Body, data)
	if message != "" {
		return newRes("fail").message(message), status
	}

	message, status = services.UpdateComment(userId, data)
	if message != "" {
		return newRes("fail").message(message), status
	}

	return newRes("success"), http.StatusOK
}

func deleteComment(req *http.Request) (res interface{}, statusCode int) {
	userId := getUserIdFromContext(req)
	commentId := getParam(req, "commentId")
	message, status := services.DeleteComment(userId, commentId)
	if message != "" {
		return newRes("fail").message(message), status
	}

	return newRes("success"), http.StatusOK
}
