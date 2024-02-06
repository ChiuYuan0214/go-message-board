package routes

import (
	"net/http"
	"security/services"
	"security/utils"
)

var updateProfileMap MethodMapType = map[string]HandlerType{}

type UpdateProfileData struct {
	Username string `json:"username"`
	Phone    string `json:"phone"`
	Job      string `json:"job"`
	Address  string `json:"address"`
}

func init() {
	updateProfileMap.post(updateProfileInfo)
}

func handleUpdateProfile(writer http.ResponseWriter, req *http.Request) {
	setContentType(writer, "json")
	res, status := updateProfileMap.useHandler(req)
	DoResponse(res, status, writer)
}

func updateProfileInfo(req *http.Request) (res interface{}, statusCode int) {
	userId := getUserIdFromContext(req)
	data := UpdateProfileData{}
	message, status := utils.ParseBody(req.Body, &data)
	if message != "" {
		return newRes("fail").message(message), status
	}

	message, status = services.UpdateColumnsById(data, &userId)
	if message != "" {
		return newRes("fail").message(message), status
	}

	return newRes("success"), http.StatusOK
}
