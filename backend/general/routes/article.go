package routes

import (
	"encoding/json"
	"general/services"
	"general/types"
	"general/utils"
	"net/http"
	"strconv"
)

var articleMap = MethodMapType{}

func init() {
	articleMap.get(getArticle).
		put(authMethod(updateArticle)).
		post(authMethod(addArticle)).
		delete(authMethod(deleteArticle))
}

func handleArticle(writer http.ResponseWriter, req *http.Request) {
	setContentType(writer, "json")
	res, status := articleMap.useHandler(req)
	DoResponse(res, status, writer)
}

func getArticle(req *http.Request) (res interface{}, statusCode int) {
	query := getQuery(req)
	id := query.Get("articleId")

	article, status := services.GetArticle(id)
	if status != 0 {
		var message string
		switch status {
		case http.StatusInternalServerError:
			message = "failed to execute query."
		case http.StatusBadRequest:
			message = "id does not exist."
		default:
			message = "something went wrong."
		}
		return newRes("fail").message(message), status
	}

	tagList := services.GetTagsByArticleId(id)
	article.Tags = tagList

	return newRes("success").setItem("data", *article), http.StatusAccepted
}

func addArticle(req *http.Request) (res interface{}, statusCode int) {
	userId := getUserIdFromContext(req)
	var data types.AddArticleData
	err := json.NewDecoder(req.Body).Decode(&data)
	if isErr(err) {
		return newRes("fail").message("body format was wrong."), http.StatusBadRequest
	}

	id := services.InsertArticle(userId, &data)
	if id == 0 {
		return newRes("fail").message("failed to insert article."), http.StatusInternalServerError
	}

	if !services.InsertTags(id, &data.Tags) {
		return newRes("fail").message("failed to insert tags."), http.StatusInternalServerError
	}

	return newRes("success").setId(id), http.StatusAccepted
}

func updateArticle(req *http.Request) (res interface{}, status int) {
	userId := getUserIdFromContext(req)
	articleId := utils.StringToInt64(getParam(req, "articleId"))
	if articleId == 0 {
		return newRes("fail").message("articleId format incorrect."), http.StatusBadRequest
	}

	var data types.UpdateArticleData
	err := json.NewDecoder(req.Body).Decode(&data)
	if isErr(err) {
		return newRes("fail").message("body format was wrong."), http.StatusBadRequest
	}

	message, status := services.UpdateArticle(userId, articleId, &data)
	if message != "" {
		return newRes("fail").message(message), status
	}
	if !services.DeleteRemovedTags(articleId, &data.Tags) || !services.InsertTags(articleId, &data.Tags) {
		return newRes("fail").message("failed to update tags."), http.StatusInternalServerError
	}

	return newRes("success"), http.StatusAccepted
}

func deleteArticle(req *http.Request) (res interface{}, statusCode int) {
	userId := getUserIdFromContext(req)
	id := getParam(req, "articleId")

	articleId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return newRes("fail").message("articleId format incorrect."), http.StatusBadRequest
	}

	message, status := services.DeleteArticle(userId, articleId)
	if message != "" {
		return newRes("fail").message(message), status
	}

	return newRes("success"), http.StatusAccepted
}
