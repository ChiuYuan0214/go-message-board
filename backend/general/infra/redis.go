package infra

import (
	"context"
	"fmt"
	"general/constants"
	"general/types"
	"log"
	"strconv"
	"time"

	"github.com/go-redis/redis/v8"
)

var _ Cache = (*Redis)(nil)

type Redis struct {
	client *redis.Client
}

func (i *Redis) Run() (err error) {
	i.client = redis.NewClient(&redis.Options{
		Addr:     constants.REDIS_IP,
		Password: constants.REDIS_PASSWORD,
		DB:       0,
	})

	_, err = i.client.Ping(context.Background()).Result()
	if err != nil {
		panic("Unable to connect to redis.")
	}
	fmt.Print("connected to redis")

	return
}

func (i *Redis) Stop() {
	i.client.Close()
}

func (i *Redis) ctx() context.Context {
	return context.Background()
}

func (i *Redis) SetToken(userId int64, token types.Token) error {
	return i.client.Set(i.ctx(), strconv.FormatInt(userId, 10), token.Token,
		time.Since(time.Unix(token.ExpireTime, 0))).Err()
}

func (i *Redis) GetToken(userId int64) (string, error) {
	return i.client.Get(i.ctx(), strconv.FormatInt(userId, 10)).Result()
}

func (i *Redis) HMSet(key string, vcm *map[string]string) error {
	return i.client.HMSet(i.ctx(), key, *vcm).Err()
}

func (i *Redis) HGetAll(key string) map[string]string {
	res, err := i.client.HGetAll(i.ctx(), key).Result()
	if err != nil {
		log.Println("failed to pull view count cache from redis.")
	}
	return res
}

func (i *Redis) HGet(key1 string, key2 string) string {
	res, err := i.client.HGet(i.ctx(), key1, key2).Result()
	if err != nil {
		return ""
	}
	return res
}

func (i *Redis) HSet(key string, mapKey string, value string) error {
	return i.client.HSet(i.ctx(), key, mapKey, value).Err()
}

func (i *Redis) ZAdd(key string, list *([]string)) error {
	zSlice := []*redis.Z{}
	for _, id := range *list {
		z := &redis.Z{
			Score:  0,
			Member: id,
		}
		zSlice = append(zSlice, z)
	}
	return i.client.ZAdd(i.ctx(), key, zSlice...).Err()
}

func (i *Redis) ZRange(key string, page, pageSize int64) []string {
	start := (page - 1) * pageSize
	end := start + pageSize - 1
	res, err := i.client.ZRange(i.ctx(), key, start, end).Result()
	if err != nil {
		return []string{}
	}
	return res
}

func (i *Redis) RPush(key string, list []string) error {
	var interfaceList []interface{}
	for _, item := range list {
		interfaceList = append(interfaceList, item)
	}
	return i.client.RPush(i.ctx(), key, interfaceList...).Err()
}

func (i *Redis) LRange(key string, page, size int64) []string {
	start := (page - 1) * size
	end := start + size - 1
	list, err := i.client.LRange(i.ctx(), key, start, end).Result()
	if err != nil {
		return []string{}
	}
	return list
}

func (i *Redis) LRangeAll(key string) []string {
	list, err := i.client.LRange(i.ctx(), key, 0, -1).Result()
	if err != nil {
		return []string{}
	}
	return list
}

func (i *Redis) SAdd(key string, val string) error {
	return i.client.SAdd(i.ctx(), key, val).Err()
}

func (i *Redis) SMembers(key string) []string {
	list, err := i.client.SMembers(i.ctx(), key).Result()
	if err != nil {
		return []string{}
	}
	return list
}

func (i *Redis) Del(key string) error {
	return i.client.Del(i.ctx(), key).Err()
}
