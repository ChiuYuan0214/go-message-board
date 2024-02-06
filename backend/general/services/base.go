package services

import (
	"database/sql"
	"general/types"
)

var connPool *sql.DB
var cache *types.RedisCache

func UsePool(db *sql.DB) {
	connPool = db
}

func UseCache(redisCache *types.RedisCache) {
	cache = redisCache
}
