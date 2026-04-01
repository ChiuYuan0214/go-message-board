package infra

import (
	"chat/types"
	"database/sql"
)

type Cache interface {
	GetToken(userId uint64) (string, error)
}

type RDB interface {
	DB() *sql.DB
}

type Dynamo interface {
	Client() *types.DynamoClient
}
