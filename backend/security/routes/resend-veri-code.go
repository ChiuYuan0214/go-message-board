package routes

import (
	"net/http"
	"security/services"
	"security/utils"
)

type ResendCodeData struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

var resendCodeMap MethodMapType = map[string]HandlerType{}

func init() {
	resendCodeMap.post(resendCode)
}

func handleResendCode(writer http.ResponseWriter, req *http.Request) {
	setContentType(writer, "json")
	res, status := resendCodeMap.useHandler(req)
	DoResponse(res, status, writer)
}

func resendCode(req *http.Request) (res interface{}, statusCode int) {
	data := &ResendCodeData{}
	message, status := utils.ParseBody(req.Body, data)
	if message != "" {
		return newRes("fail").message(message), status
	}

	if data.Email == "" || data.Password == "" {
		return newRes("fail").message("email and password cannot be empty."), http.StatusBadRequest
	}

	userId := services.VerifyPasswordByEmail(&data.Email, &data.Password)
	if userId == 0 {
		return newRes("fail").message("password incorrect"), http.StatusBadRequest
	}
	if userId == -1 {
		return newRes("fail").message("account already active."), http.StatusOK
	}

	veriCode := utils.GenerateCode()
	isSent := utils.SendVerifyCode(data.Email, veriCode.Code)
	if !isSent || !services.InvalidateVerificationCodes(userId) {
		return newRes("fail").message("failed to send code."), http.StatusInternalServerError
	}

	codeId := services.InsertVerificationCode(userId, veriCode.Code, veriCode.ExpireTime)
	if codeId == 0 {
		return newRes("fail").message("failed to record verification code."), http.StatusInternalServerError
	}
	services.ScheduleCodeInvalidation(codeId, veriCode)

	return newRes("success"), http.StatusOK
}
