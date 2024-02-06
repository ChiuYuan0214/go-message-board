package routes

import (
	"net/http"
	"security/services"
	"security/utils"
)

type UpdatePassData struct {
	OldPassword string `json:"oldPassword"`
	NewPassword string `json:"newPassword"`
}

var updatePassMap MethodMapType = map[string]HandlerType{}

func init() {
	updatePassMap.put(updatePassword)
}

func handleUpdatePassword(writer http.ResponseWriter, req *http.Request) {
	setContentType(writer, "json")
	res, status := updatePassMap.useHandler(req)
	DoResponse(res, status, writer)
}

func updatePassword(req *http.Request) (res interface{}, statusCode int) {
	userId := getUserIdFromContext(req)
	data := &UpdatePassData{}
	message, status := utils.ParseBody(req.Body, data)
	if message != "" {
		return newRes("fail").message(message), status
	}

	if userId == 0 || data.OldPassword == "" || data.NewPassword == "" {
		return newRes("fail").message("userId, old password and new password cannot be empty"), http.StatusBadRequest
	}

	if !services.VerifyPasswordByUserId(&userId, &data.OldPassword) {
		return newRes("fail").message("userId or old password incorrect."), http.StatusOK
	}

	if !services.UpdatePassword(&userId, &data.NewPassword) {
		return newRes("fail").message("failed to update password."), http.StatusInternalServerError
	}

	return newRes("success"), http.StatusOK
}
