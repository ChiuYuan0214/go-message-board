package repo

import (
	"security/entities"
	"security/infra"
	"security/types"

	"gorm.io/gorm"
)

var _ Auth = (*AuthImpl)(nil)

type AuthImpl struct {
	db    infra.RDB
	cache infra.Cache
}

func NewAuth(db infra.RDB, cache infra.Cache) *AuthImpl {
	return &AuthImpl{
		db:    db,
		cache: cache,
	}
}

func (r *AuthImpl) GetLoginCredentialByEmail(email string) (userId uint64, hashedPassword string, err error) {
	var user entities.User
	err = r.db.Orm().
		Select("user_id", "password").
		Where("email = ?", email).
		First(&user).Error
	if err == gorm.ErrRecordNotFound {
		err = nil
		return
	}

	userId = user.UserId
	hashedPassword = user.Password
	return
}

func (r *AuthImpl) SetToken(userId uint64, token types.Token) error {
	return r.cache.SetToken(userId, token)
}

func (r *AuthImpl) GetToken(userId uint64) (string, error) {
	return r.cache.GetToken(userId)
}
