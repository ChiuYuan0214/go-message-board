package routes

import (
	"net/http"
	"security/services"
	"strings"
)

type UsersHandler struct {
	router      Router
	authService services.Auth
	userService services.User
	usersMap    MethodMapType
}

func NewUsersHandler(router Router, authService services.Auth, userService services.User) *UsersHandler {
	return &UsersHandler{
		router:      router,
		authService: authService,
		userService: userService,
	}
}

func (h *UsersHandler) Run() {
	h.usersMap = make(MethodMapType)
	h.usersMap.get(h.getUsers)
	h.router.Handle("/users", authMiddle(h.authService, h.handleUsers))
}

func (h *UsersHandler) handleUsers(writer http.ResponseWriter, req *http.Request) {
	setHeader(writer, "json")
	res, status := h.usersMap.useHandler(writer, req)
	DoResponse(res, status, writer)
}

func (h *UsersHandler) getUsers(req *http.Request) (res interface{}, statusCode int) {
	name := getParam(req, "name")
	userId := getUserIdFromContext(req)
	filteredUsers := h.userService.GetUsers(strings.TrimSpace(name), userId)
	return newRes("success").setList("users", filteredUsers), http.StatusOK
}
