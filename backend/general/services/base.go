package services

import (
	"general/types"

	"gorm.io/gorm"
)

var db *gorm.DB
var cache *types.RedisCache

func UsePool(DB *gorm.DB) {
	db = DB
}

func UseCache(redisCache *types.RedisCache) {
	cache = redisCache
}
