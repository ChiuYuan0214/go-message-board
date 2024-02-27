package routes

import (
	"encoding/json"
	"general/services"
	"general/types"
	"net/http"
)

var followMap = MethodMapType{}

func init() {
	followMap.post(addFollow).delete(removeFollow)
}

func handleFollow(writer http.ResponseWriter, req *http.Request) {
	setHeader(writer, "json")
	res, status := followMap.useHandler(writer, req)
	DoResponse(res, status, writer)
}

func addFollow(req *http.Request) (res interface{}, statusCode int) {
	var data types.FollowData
	err := json.NewDecoder(req.Body).Decode(&data)
	userId := getUserIdFromContext(req)
	if err != nil {
		return newRes("fail").message("body format was wrong."), http.StatusBadRequest
	}
	if userId == 0 || data.Followee == 0 {
		return newRes("fail").message("userId and followee cannot be empty"), http.StatusBadRequest
	}
	if !services.AddFollow(userId, data.Followee) {
		return newRes("fail").message("failed to insert data"), http.StatusBadRequest
	}
	return newRes("success"), http.StatusOK
}

func removeFollow(req *http.Request) (res interface{}, statusCode int) {
	var data types.FollowData
	err := json.NewDecoder(req.Body).Decode(&data)
	userId := getUserIdFromContext(req)
	if err != nil {
		return newRes("fail").message("body format was wrong."), http.StatusBadRequest
	}
	if userId == 0 || data.Followee == 0 {
		return newRes("fail").message("userId and followee cannot be empty"), http.StatusBadRequest
	}
	if !services.RemoveFollow(userId, data.Followee) {
		return newRes("fail").message("failed to delete data"), http.StatusBadRequest
	}
	return newRes("success"), http.StatusOK
}
