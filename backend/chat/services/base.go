package services

import (
	"chat/store"
	"chat/types"
	"database/sql"
)

var connPool *sql.DB
var cache *types.RedisCache
var mongo *types.MongoClient
var chatStore = store.GetChatStore()

func UsePool(db *sql.DB) {
	connPool = db
}

func UseCache(redisCache *types.RedisCache) {
	cache = redisCache
}

func UseMongo(mc *types.MongoClient) {
	mongo = mc
}
