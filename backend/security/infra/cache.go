package infra

import (
	"context"
	"fmt"
	"security/constants"
	"security/types"
	"strconv"
	"time"

	"github.com/go-redis/redis/v8"
)

var _ Cache = (*Redis)(nil)

type Redis struct {
	client *redis.Client
	ctx    context.Context
}

func (r *Redis) Run() (err error) {
	r.client = redis.NewClient(&redis.Options{
		Addr:     constants.REDIS_IP,
		Password: constants.REDIS_PASSWORD,
		DB:       0,
	})

	r.ctx = context.Background()
	_, err = r.client.Ping(r.ctx).Result()
	if err != nil {
		return err
	}
	fmt.Print("connected to redis")
	return
}

func (r *Redis) Stop() {
	if r.client != nil {
		r.client.Close()
	}
}

func (r *Redis) SetToken(userId uint64, token types.Token) error {
	return r.client.Set(r.ctx, strconv.FormatUint(userId, 10), token.Token,
		time.Until(time.Unix(token.ExpireTime, 0))).Err()
}

func (r *Redis) GetToken(userId uint64) (string, error) {
	return r.client.Get(r.ctx, strconv.FormatUint(userId, 10)).Result()
}
