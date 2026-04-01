package routes

import (
	"fmt"
	"net/http"
	"security/utils"
)

type VerifyData struct {
	UserId uint64 `json:"userId"`
	Code   int32  `json:"code"`
}

func (h *RegisterHandler) handleVerifyCode(writer http.ResponseWriter, req *http.Request) {
	setHeader(writer, "json")
	res, status := h.verifyMap.useHandler(writer, req)
	DoResponse(res, status, writer)
}

func (h *RegisterHandler) doVerify(req *http.Request) (res interface{}, statusCode int) {
	data := &VerifyData{}
	message, status := utils.ParseBody(req.Body, data)
	if message != "" {
		return newRes("fail").message(message), status
	}
	if data.Code == 0 || data.UserId == 0 {
		return newRes("fail").message("userId and code cannot be empty."), http.StatusBadRequest
	}

	code, err := h.registerService.GetActiveVerificationCode(data.UserId)
	if isErr(err) {
		return newRes("fail").message("failed to query."), http.StatusInternalServerError
	}
	if code == "" {
		return newRes("fail").message("code does not exist."), http.StatusOK
	}

	if code != fmt.Sprintf("%06d", data.Code) {
		return newRes("fail").message("code does not match."), http.StatusOK
	}

	h.registerService.ActivateUser(data.UserId)

	// generate token
	token := h.authService.GenerateToken(data.UserId)
	if token == nil {
		return newRes("fail").message("failed to generate token."), http.StatusInternalServerError
	}

	return newRes("success").setItem("token", token.Token).setItem("expireTime", token.ExpireTime), http.StatusOK
}
