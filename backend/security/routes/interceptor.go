package routes

import (
	"context"
	"net/http"
	"security/services"
	"strings"
)

func authMiddle(next http.HandlerFunc) http.HandlerFunc {
	return func(writer http.ResponseWriter, req *http.Request) {
		if req.Method == http.MethodOptions {
			next.ServeHTTP(writer, req)
			return
		}

		authHeader := req.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(writer, "Unauthorized", http.StatusUnauthorized)
			return
		}

		tokenSlice := strings.Split(authHeader, " ")
		if len(tokenSlice) != 2 || tokenSlice[0] != "Bearer" {
			http.Error(writer, "Invalid token format", http.StatusUnauthorized)
			return
		}

		token := tokenSlice[1]
		userId := services.GetUserIdFromToken(token)
		if userId == 0 {
			http.Error(writer, "Unauthorized", http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(req.Context(), "userId", userId)
		next.ServeHTTP(writer, req.WithContext(ctx))
	}
}
