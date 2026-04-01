package infra

import (
	"security/types"

	"gorm.io/gorm"
)

type RDB interface {
	Orm() *gorm.DB
}

type Cache interface {
	SetToken(userId uint64, token types.Token) error
	GetToken(userId uint64) (string, error)
}
