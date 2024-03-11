package routes

import (
	"net/http"
	"security/store"
	"strings"
)

var usersMap MethodMapType = map[string]HandlerType{}

func init() {
	usersMap.get(getUsers)
}

func handleUsers(writer http.ResponseWriter, req *http.Request) {
	setHeader(writer, "json")
	res, status := usersMap.useHandler(writer, req)
	DoResponse(res, status, writer)
}

func getUsers(req *http.Request) (res interface{}, statusCode int) {
	name := getParam(req, "name")
	userId := getUserIdFromContext(req)
	users := store.GetUsers()
	filteredUsers := []store.User{}

	for _, u := range users {
		if strings.Contains(u.UserName, name) && userId != u.UserId {
			filteredUsers = append(filteredUsers, u)
		}
	}
	return newRes("success").setList("users", filteredUsers), http.StatusOK
}
