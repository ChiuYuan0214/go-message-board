package routes

import (
	"encoding/json"
	"general/services"
	"general/types"
	"log"
	"net/http"
)

var collectionMap = MethodMapType{}

func init() {
	collectionMap.get(getCollections).post(addCollection).delete(removeCollection)
}

func handleCollection(writer http.ResponseWriter, req *http.Request) {
	setHeader(writer, "json")
	res, status := collectionMap.useHandler(writer, req)
	DoResponse(res, status, writer)
}

func getCollections(req *http.Request) (res interface{}, statusCode int) {
	userId := getUserIdFromContext(req)
	page, size := getPageSize(req)
	if userId == 0 {
		return newRes("fail").message("userId not valid."), http.StatusBadRequest
	}
	data := services.GetCollections(userId, page, size)
	return newRes("success").setList("list", data), http.StatusOK
}

func addCollection(req *http.Request) (res interface{}, statusCode int) {
	userId := getUserIdFromContext(req)

	var body types.WriteCollectionData
	err := json.NewDecoder(req.Body).Decode(&body)
	if err != nil {
		log.Println(err)
		return newRes("fail").message("body format was wrong."), http.StatusBadRequest
	}
	if userId == 0 || body.ArticleId == 0 {
		return newRes("fail").message("userId and articleId cannot be empty."), http.StatusBadRequest
	}
	if !services.AddCollection(userId, body.ArticleId) {
		return newRes("fail").message("failed to add collection."), http.StatusInternalServerError
	}
	return newRes("success"), http.StatusOK
}

func removeCollection(req *http.Request) (res interface{}, statusCode int) {
	userId := getUserIdFromContext(req)
	var body types.WriteCollectionData
	err := json.NewDecoder(req.Body).Decode(&body)
	if err != nil {
		return newRes("fail").message("body format was wrong."), http.StatusBadRequest
	}
	if userId == 0 || body.ArticleId == 0 {
		return newRes("fail").message("userId and articleId cannot be empty."), http.StatusBadRequest
	}
	if !services.RemoveCollection(userId, body.ArticleId) {
		return newRes("fail").message("failed to remove collection."), http.StatusInternalServerError
	}
	return newRes("success"), http.StatusOK
}
