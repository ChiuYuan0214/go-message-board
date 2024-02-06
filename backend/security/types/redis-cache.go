package types

import (
	"context"
	"strconv"
	"time"

	"github.com/go-redis/redis/v8"
)

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
