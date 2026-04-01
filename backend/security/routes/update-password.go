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

type ProfileHandler struct {
	router         Router
	authService    services.Auth
	profileService services.Profile
	updatePassMap  MethodMapType
	updateProfMap  MethodMapType
	uploadImageMap MethodMapType
}

func NewProfileHandler(router Router, authService services.Auth, profileService services.Profile) *ProfileHandler {
	return &ProfileHandler{
		router:         router,
		authService:    authService,
		profileService: profileService,
	}
}

func (h *ProfileHandler) Run() {
	h.updatePassMap = make(MethodMapType)
	h.updateProfMap = make(MethodMapType)
	h.uploadImageMap = make(MethodMapType)

	h.updatePassMap.put(h.updatePassword)
	h.updateProfMap.post(h.updateProfileInfo)
	h.uploadImageMap.post(h.upload)

	h.router.Handle("/updatePassword", authMiddle(h.authService, h.handleUpdatePassword))
	h.router.Handle("/updateProfile", authMiddle(h.authService, h.handleUpdateProfile))
	h.router.Handle("/uploadImage", authMiddle(h.authService, h.handleUploadImage))
}

func (h *ProfileHandler) handleUpdatePassword(writer http.ResponseWriter, req *http.Request) {
	setHeader(writer, "json")
	res, status := h.updatePassMap.useHandler(writer, req)
	DoResponse(res, status, writer)
}

func (h *ProfileHandler) updatePassword(req *http.Request) (res interface{}, statusCode int) {
	userId := getUserIdFromContext(req)
	data := &UpdatePassData{}
	message, status := utils.ParseBody(req.Body, data)
	if message != "" {
		return newRes("fail").message(message), status
	}

	if userId == 0 || data.OldPassword == "" || data.NewPassword == "" {
		return newRes("fail").message("userId, old password and new password cannot be empty"), http.StatusBadRequest
	}

	if !h.profileService.VerifyPasswordByUserId(&userId, &data.OldPassword) {
		return newRes("fail").message("userId or old password incorrect."), http.StatusOK
	}

	if !h.profileService.UpdatePassword(&userId, &data.NewPassword) {
		return newRes("fail").message("failed to update password."), http.StatusInternalServerError
	}

	return newRes("success"), http.StatusOK
}
