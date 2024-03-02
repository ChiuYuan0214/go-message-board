package services

import (
	"chat/store"
	"chat/types"
	"database/sql"
)

var connPool *sql.DB
var cache *types.RedisCache

var dynamo *types.DynamoClient
var chatStore = store.GetChatStore()

func UsePool(db *sql.DB) {
	connPool = db
}

func UseCache(redisCache *types.RedisCache) {
	cache = redisCache
}

func UseDynamo(dc *types.DynamoClient) {
	dynamo = dc
}
