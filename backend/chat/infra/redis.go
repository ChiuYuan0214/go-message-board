package infra

import (
	"chat/constants"
	"context"
	"fmt"
	"strconv"

	"github.com/go-redis/redis/v8"
)

var _ Cache = (*Redis)(nil)

type Redis struct {
	client *redis.Client
}

func (r *Redis) Run() (err error) {
	r.client = redis.NewClient(&redis.Options{
		Addr:     constants.REDIS_IP,
		Password: constants.REDIS_PASSWORD,
		DB:       0,
	})
	_, err = r.client.Ping(r.ctx()).Result()
	if err != nil {
		return err
	}
	fmt.Println("connected to redis")
	return
}

func (r *Redis) Stop() {
	if r.client != nil {
		r.client.Close()
	}
}

func (r *Redis) ctx() context.Context {
	return context.Background()
}

func (r *Redis) GetToken(userId uint64) (string, error) {
	return r.client.Get(r.ctx(), strconv.FormatUint(userId, 10)).Result()
}
