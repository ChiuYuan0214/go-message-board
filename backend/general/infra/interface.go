package infra

import (
	"general/types"

	"gorm.io/gorm"
)

type RDB interface {
	Orm() *gorm.DB
}

type Cache interface {
	SetToken(userId int64, token types.Token) error
	GetToken(userId int64) (string, error)
	HMSet(key string, vcm *map[string]string) error
	HGetAll(key string) map[string]string
	HGet(key1 string, key2 string) string
	HSet(key string, mapKey string, value string) error
	ZAdd(key string, list *([]string)) error
	ZRange(key string, page, pageSize int64) []string
	RPush(key string, list []string) error
	LRange(key string, page, size int64) []string
	LRangeAll(key string) []string
	SAdd(key string, val string) error
	SMembers(key string) []string
	Del(key string) error
}
