package utils

import (
	"encoding/json"
	"general/constants"
	"io"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
)

func IsAuth(req *http.Request) uint64 {
	authHeader := req.Header.Get("Authorization")
	if authHeader == "" {
		return 0
	}

	tokenSlice := strings.Split(authHeader, " ")
	if len(tokenSlice) != 2 || tokenSlice[0] != "Bearer" {
		return 0
	}

	token := tokenSlice[1]
	userId := GetUserIdFromToken(token)
	if userId == 0 {
		return 0
	}
	return userId
}

func ParseBody(body io.ReadCloser, containerAddr interface{}) (string, int) {
	decoder := json.NewDecoder(body)
	defer body.Close()

	if !decoder.More() {
		return "body was empty.", http.StatusBadRequest
	}

	err := decoder.Decode(containerAddr)
	if err != nil {
		return "body format was wrong.", http.StatusBadRequest
	}

	return "", 0
}

func GetUserIdFromToken(srcToken string) uint64 {
	token, err := jwt.Parse(srcToken, func(t *jwt.Token) (interface{}, error) {
		return []byte(constants.JWT_HS256_SECRET_KEY), nil
	})

	if err != nil || !token.Valid {
		return 0
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return 0
	}

	userId, ok := claims["sub"].(float64)
	if !ok {
		return 0
	}

	return uint64(userId)
}
