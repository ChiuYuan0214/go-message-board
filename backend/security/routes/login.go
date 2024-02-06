package routes

import (
	"net/http"
	"security/services"
	"security/utils"
)

type LoginData struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type RefreshData struct {
	UserId int64  `json:"userId"`
	Token  string `json:"token"`
}

var loginMap MethodMapType = map[string]HandlerType{}

func init() {
	loginMap.post(login).put(refreshToken)
}

func handleLogin(writer http.ResponseWriter, req *http.Request) {
	setContentType(writer, "json")
	res, status := loginMap.useHandler(req)
	DoResponse(res, status, writer)
}

func login(req *http.Request) (res interface{}, statusCode int) {
	data := &LoginData{}
	message, status := utils.ParseBody(req.Body, data)
	if message != "" {
		return newRes("fail").message(message), status
	}

	userId, token := services.Login(data.Email, data.Password)
	if userId == 0 {
		return newRes("fail").message("account not exist"), http.StatusOK
	}
	if token == nil {
		return newRes("fail").message("failed to create token"), http.StatusInternalServerError
	}

	return newRes("success").setItem("token", token.Token).setItem("expireTime", token.ExpireTime), http.StatusOK
}

func refreshToken(req *http.Request) (res interface{}, statusCode int) {
	data := &RefreshData{}
	message, status := utils.ParseBody(req.Body, data)
	if message != "" {
		return newRes("fail").message(message), status
	}

	if data.UserId == 0 || data.Token == "" {
		return newRes("fail").message("userId and token cannot be empty"), http.StatusBadRequest
	}

	if !services.VerifyToken(data.UserId, data.Token) {
		return newRes("fail").message("token was incorrect."), http.StatusBadRequest
	}

	token := services.GenerateToken(data.UserId)

	return newRes("success").setItem("token", token.Token).setItem("expireTime", token.ExpireTime), http.StatusOK
}
