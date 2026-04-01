package repo

import (
	"database/sql"
	"security/infra"
	"security/types"
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
	err = r.db.DB().QueryRow("select user_id, password from users where email = ?", email).Scan(&userId, &hashedPassword)
	if err == sql.ErrNoRows {
		err = nil
	}
	return
}

func (r *AuthImpl) SetToken(userId uint64, token types.Token) error {
	return r.cache.SetToken(userId, token)
}

func (r *AuthImpl) GetToken(userId uint64) (string, error) {
	return r.cache.GetToken(userId)
}
