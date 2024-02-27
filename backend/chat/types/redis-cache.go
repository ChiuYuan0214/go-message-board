package types

import (
	"context"
	"strconv"

	"github.com/go-redis/redis/v8"
)

type RedisCache struct {
	Client *redis.Client
	Ctx    context.Context
}

func (rc *RedisCache) GetToken(userId uint64) (string, error) {
	return rc.Client.Get(rc.Ctx, strconv.FormatUint(userId, 10)).Result()
}
