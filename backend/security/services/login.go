package services

import (
	"security/constants"
	"security/types"
	"security/utils"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func Login(email string, password string) (userId int64, token *types.Token) {
	var hashedPassword string
	row := connPool.QueryRow("select user_id, password from users where email = ?", email)
	row.Scan(&userId, &hashedPassword)
	matched := utils.VerifyPassword(&hashedPassword, &password)
	if !matched {
		return 0, nil
	}
	token = GenerateToken(userId)
	return userId, token
}

func GenerateToken(userId int64) *types.Token {
	expireTime := time.Now().Add(30 * time.Minute).Unix()
	claims := jwt.MapClaims{
		"sub": userId,
		"exp": expireTime,
	}

	tokenSource := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := tokenSource.SignedString([]byte(constants.JWT_HS256_SECRET_KEY))
	if err != nil {
		return nil
	}
	token := types.Token{Token: tokenString, ExpireTime: expireTime}
	cache.SetToken(userId, token)

	return &token
}

func VerifyToken(userId int64, token string) bool {
	actualToken, err := cache.GetToken(userId)
	if err != nil || token != actualToken {
		return false
	}
	return true
}

func GetUserIdFromToken(srcToken string) int64 {
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

	return int64(userId)
}
