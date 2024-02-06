package types

import (
	"context"
	"log"
	"strconv"
	"time"

	"github.com/go-redis/redis/v8"
)

// type HashMap map[string]string

// func (i HashMap) MarshalBinary() (data []byte, err error) {
// 	bytes, err := json.Marshal(i)
// 	return bytes, err
// }

type RedisCache struct {
	Client *redis.Client
	Ctx    context.Context
}

func (rc *RedisCache) SetToken(userId int64, token Token) error {
	return rc.Client.Set(rc.Ctx, strconv.FormatInt(userId, 10), token.Token,
		time.Since(time.Unix(token.ExpireTime, 0))).Err()
}

func (rc *RedisCache) GetToken(userId int64) (string, error) {
	return rc.Client.Get(rc.Ctx, strconv.FormatInt(userId, 10)).Result()
}

func (rc *RedisCache) HMSet(key string, vcm *map[string]string) error {
	return rc.Client.HMSet(rc.Ctx, key, *vcm).Err()
}

func (rc *RedisCache) HGetAll(key string) map[string]string {
	res, err := rc.Client.HGetAll(rc.Ctx, key).Result()
	if err != nil {
		log.Println("failed to pull view count cache from redis.")
	}
	return res
}

func (rc *RedisCache) HGet(key1 string, key2 string) string {
	res, err := rc.Client.HGet(rc.Ctx, key1, key2).Result()
	if err != nil {
		return ""
	}
	return res
}

func (rc *RedisCache) HSet(key string, mapKey string, value string) error {
	return rc.Client.HSet(rc.Ctx, key, mapKey, value).Err()
}

func (rc *RedisCache) ZAdd(key string, list *([]string)) error {
	zSlice := []*redis.Z{}
	for _, id := range *list {
		z := &redis.Z{
			Score:  0,
			Member: id,
		}
		zSlice = append(zSlice, z)
	}
	return rc.Client.ZAdd(rc.Ctx, key, zSlice...).Err()
}

func (rc *RedisCache) ZRange(key string, page, pageSize int64) []string {
	start := (page - 1) * pageSize
	end := start + pageSize - 1
	res, err := rc.Client.ZRange(rc.Ctx, key, start, end).Result()
	if err != nil {
		return []string{}
	}
	return res
}

func (rc *RedisCache) RPush(key string, list *([]string)) error {
	var interfaceList []interface{}
	for _, item := range *list {
		interfaceList = append(interfaceList, item)
	}
	return rc.Client.RPush(rc.Ctx, key, interfaceList...).Err()
}

func (rc *RedisCache) LRange(key string, page, size int64) []string {
	start := (page - 1) * size
	end := start + size - 1
	list, err := rc.Client.LRange(rc.Ctx, key, start, end).Result()
	if err != nil {
		return []string{}
	}
	return list
}

func (rc *RedisCache) LRangeAll(key string) []string {
	list, err := rc.Client.LRange(rc.Ctx, key, 0, -1).Result()
	if err != nil {
		return []string{}
	}
	return list
}

func (rc *RedisCache) SAdd(key string, val string) error {
	return rc.Client.SAdd(rc.Ctx, key, val).Err()
}

func (rc *RedisCache) SMembers(key string) []string {
	list, err := rc.Client.SMembers(rc.Ctx, key).Result()
	if err != nil {
		return []string{}
	}
	return list
}

func (rc *RedisCache) Del(key string) error {
	return rc.Client.Del(rc.Ctx, key).Err()
}
