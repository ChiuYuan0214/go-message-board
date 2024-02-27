package routes

import (
	"general/services"
	"general/utils"
	"net/http"
)

var profileMap = MethodMapType{}

func init() {
	profileMap.get(getProfile)
}

func handleProfile(writer http.ResponseWriter, req *http.Request) {
	setHeader(writer, "json")
	res, status := profileMap.useHandler(writer, req)
	DoResponse(res, status, writer)
}

func getProfile(req *http.Request) (res interface{}, statusCode int) {
	userId := getParam(req, "userId")

	if userId != "" {
		profile, status := services.GetProfileWithId(userId)
		if profile != nil {
			return newRes("success").setItem("data", *profile), status
		}
		if status == http.StatusBadRequest {
			return newRes("fail").message("invalid input."), status
		}
		return newRes("fail").message("something went wrong."), status
	}

	id := utils.IsAuth(req)
	profile, status := services.GetProfileWithToken(id)
	if profile != nil {
		return newRes("success").setItem("data", *profile), status
	}
	return newRes("fail").message("something went wrong."), status
}
