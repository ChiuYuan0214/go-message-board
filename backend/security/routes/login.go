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
	UserId uint64 `json:"userId"`
	Token  string `json:"token"`
}

type LoginHandler struct {
	router      Router
	authService services.Auth
	loginMap    MethodMapType
}

func NewLoginHandler(router Router, authService services.Auth) *LoginHandler {
	return &LoginHandler{
		router:      router,
		authService: authService,
	}
}

func (h *LoginHandler) Run() {
	h.loginMap = make(MethodMapType)
	h.loginMap.post(h.login).put(h.refreshToken)
	h.router.Handle("/login", h.handleLogin)
}

func (h *LoginHandler) handleLogin(writer http.ResponseWriter, req *http.Request) {
	setHeader(writer, "json")
	res, status := h.loginMap.useHandler(writer, req)
	DoResponse(res, status, writer)
}

func (h *LoginHandler) login(req *http.Request) (res interface{}, statusCode int) {
	data := &LoginData{}
	message, status := utils.ParseBody(req.Body, data)
	if message != "" {
		return newRes("fail").message(message), status
	}

	userId, token := h.authService.Login(data.Email, data.Password)
	if userId == 0 {
		return newRes("fail").message("account not exist"), http.StatusOK
	}
	if token == nil {
		return newRes("fail").message("failed to create token"), http.StatusInternalServerError
	}

	return newRes("success").setItem("userId", userId).setItem("token", token.Token).setItem("expireTime", token.ExpireTime), http.StatusOK
}

func (h *LoginHandler) refreshToken(req *http.Request) (res interface{}, statusCode int) {
	data := &RefreshData{}
	message, status := utils.ParseBody(req.Body, data)
	if message != "" {
		return newRes("fail").message(message), status
	}

	if data.UserId == 0 || data.Token == "" {
		return newRes("fail").message("userId and token cannot be empty"), http.StatusBadRequest
	}

	if !h.authService.VerifyToken(data.UserId, data.Token) {
		return newRes("fail").message("token was incorrect."), http.StatusBadRequest
	}

	token := h.authService.GenerateToken(data.UserId)
	return newRes("success").setItem("token", token.Token).setItem("expireTime", token.ExpireTime), http.StatusOK
}
