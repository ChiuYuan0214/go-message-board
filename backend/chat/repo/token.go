package repo

import "chat/infra"

var _ Token = (*TokenImpl)(nil)

type TokenImpl struct {
	cache infra.Cache
}

func (r *TokenImpl) Run() (err error) {
	return
}

func (r *TokenImpl) Stop() {}

func (r *TokenImpl) GetToken(userId uint64) (string, error) {
	return r.cache.GetToken(userId)
}
