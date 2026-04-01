package routes

import (
	"net/http"
	"security/utils"
)

type ResendCodeData struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (h *RegisterHandler) handleResendCode(writer http.ResponseWriter, req *http.Request) {
	setHeader(writer, "json")
	res, status := h.resendCodeMap.useHandler(writer, req)
	DoResponse(res, status, writer)
}

func (h *RegisterHandler) resendCode(req *http.Request) (res interface{}, statusCode int) {
	data := &ResendCodeData{}
	message, status := utils.ParseBody(req.Body, data)
	if message != "" {
		return newRes("fail").message(message), status
	}

	if data.Email == "" || data.Password == "" {
		return newRes("fail").message("email and password cannot be empty."), http.StatusBadRequest
	}

	userId := h.registerService.VerifyPasswordByEmail(&data.Email, &data.Password)
	if userId == 0 {
		return newRes("fail").message("password incorrect"), http.StatusBadRequest
	}
	if userId == -1 {
		return newRes("fail").message("account already active."), http.StatusOK
	}

	veriCode := utils.GenerateCode()
	isSent := utils.SendVerifyCode(data.Email, veriCode.Code)
	if !isSent || !h.registerService.InvalidateVerificationCodes(userId) {
		return newRes("fail").message("failed to send code."), http.StatusInternalServerError
	}

	codeId := h.registerService.InsertVerificationCode(userId, veriCode.Code, veriCode.ExpireTime)
	if codeId == 0 {
		return newRes("fail").message("failed to record verification code."), http.StatusInternalServerError
	}
	h.registerService.ScheduleCodeInvalidation(codeId, veriCode)

	return newRes("success").setItem("expireTime", veriCode.ExpireTime), http.StatusOK
}
