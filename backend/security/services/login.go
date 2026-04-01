package services

import (
	"security/constants"
	"security/repo"
	"security/types"
	"security/utils"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var _ Auth = (*AuthImpl)(nil)

type AuthImpl struct {
	authRepo repo.Auth
}

func NewAuth(authRepo repo.Auth) *AuthImpl {
	return &AuthImpl{
		authRepo: authRepo,
	}
}

func (s *AuthImpl) Login(email string, password string) (userId uint64, token *types.Token) {
	var hashedPassword string
	userId, hashedPassword, _ = s.authRepo.GetLoginCredentialByEmail(email)
	matched := utils.VerifyPassword(&hashedPassword, &password)
	if !matched {
		return 0, nil
	}
	token = s.GenerateToken(userId)
	return userId, token
}

func (s *AuthImpl) GenerateToken(userId uint64) *types.Token {
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
	if err = s.authRepo.SetToken(userId, token); err != nil {
		return nil
	}

	return &token
}

func (s *AuthImpl) VerifyToken(userId uint64, token string) bool {
	actualToken, err := s.authRepo.GetToken(userId)
	if err != nil || token != actualToken {
		return false
	}
	return true
}

func (s *AuthImpl) GetUserIdFromToken(srcToken string) uint64 {
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
