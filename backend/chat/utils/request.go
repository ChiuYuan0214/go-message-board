package utils

import (
	"chat/constants"
	"net/http"

	"github.com/dgrijalva/jwt-go"
)

func GetTokenFromQuery(r *http.Request) string {
	return r.URL.Query().Get("token")
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
