package routes

import (
	"context"
	"general/utils"
	"net/http"
)

var errRes = newRes("fail").message("Unauthorized")

func authMiddle(next http.HandlerFunc) http.HandlerFunc {
	return func(writer http.ResponseWriter, req *http.Request) {
		if req.Method == http.MethodOptions {
			next.ServeHTTP(writer, req)
			return
		}
		userId := utils.IsAuth(req)
		if userId == 0 {
			http.Error(writer, "Unauthorized", http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(req.Context(), "userId", userId)
		next.ServeHTTP(writer, req.WithContext(ctx))
	}
}

func authMethod(handler HandlerType) HandlerType {
	return func(req *http.Request) (interface{}, int) {
		userId := utils.IsAuth(req)
		if userId == 0 {
			return errRes, http.StatusUnauthorized
		}
		ctx := context.WithValue(req.Context(), "userId", userId)
		return handler(req.WithContext(ctx))
	}
}
