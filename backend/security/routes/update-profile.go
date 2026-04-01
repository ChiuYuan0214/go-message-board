package routes

import (
	"net/http"
	"security/utils"
)

type UpdateProfileData struct {
	Username string `json:"username"`
	Phone    string `json:"phone"`
	Job      string `json:"job"`
	Address  string `json:"address"`
}

func (h *ProfileHandler) handleUpdateProfile(writer http.ResponseWriter, req *http.Request) {
	setHeader(writer, "json")
	res, status := h.updateProfMap.useHandler(writer, req)
	DoResponse(res, status, writer)
}

func (h *ProfileHandler) updateProfileInfo(req *http.Request) (res interface{}, statusCode int) {
	userId := getUserIdFromContext(req)
	data := UpdateProfileData{}
	message, status := utils.ParseBody(req.Body, &data)
	if message != "" {
		return newRes("fail").message(message), status
	}

	message, status = h.profileService.UpdateColumnsById(data, &userId)
	if message != "" {
		return newRes("fail").message(message), status
	}

	return newRes("success"), http.StatusOK
}
