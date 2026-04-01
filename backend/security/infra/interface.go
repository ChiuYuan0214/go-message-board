package infra

import (
	"database/sql"
	"security/types"
)

type RDB interface {
	DB() *sql.DB
}

type Cache interface {
	SetToken(userId uint64, token types.Token) error
	GetToken(userId uint64) (string, error)
}
