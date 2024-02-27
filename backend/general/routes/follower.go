package routes

import (
	"encoding/json"
	"general/services"
	"general/types"
	"net/http"
)

var followerMap = MethodMapType{}

func init() {
	followerMap.delete(authMethod(removeFollower)).get(getFollowers)
}

func handleFollower(writer http.ResponseWriter, req *http.Request) {
	setHeader(writer, "json")
	res, status := followerMap.useHandler(writer, req)
	DoResponse(res, status, writer)
}

func removeFollower(req *http.Request) (res interface{}, statusCode int) {
	var data types.FollowerData
	userId := getUserIdFromContext(req)
	err := json.NewDecoder(req.Body).Decode(&data)
	if err != nil {
		return newRes("fail").message("body format was wrong."), http.StatusBadRequest
	}
	if userId == 0 || data.Follower == 0 {
		return newRes("fail").message("userId and follower cannot be empty"), http.StatusBadRequest
	}
	if !services.RemoveFollower(userId, data.Follower) {
		return newRes("fail").message("failed to delete data"), http.StatusBadRequest
	}
	return newRes("success"), http.StatusOK
}

func getFollowers(req *http.Request) (res interface{}, statusCode int) {
	userId := getUserIdFromQuery(req)
	if userId == 0 {
		return newRes("fail").message("userId not valid."), http.StatusBadRequest
	}
	data := services.GetFollowers(userId)
	return newRes("success").setList("list", *data), http.StatusOK
}
