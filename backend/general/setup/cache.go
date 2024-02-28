package setup

import (
	"context"
	"fmt"
	"general/constants"
	"general/types"

	"github.com/go-redis/redis/v8"
)

func InitCache() *types.RedisCache {
	client := redis.NewClient(&redis.Options{
		Addr:     constants.REDIS_IP,
		Password: constants.REDIS_PASSWORD,
		DB:       0,
	})

	var ctx = context.Background()
	_, err := client.Ping(ctx).Result()
	if err != nil {
		panic("Unable to connect to redis.")
	}
	fmt.Print("connected to redis")

	return &types.RedisCache{Client: client, Ctx: ctx}
}
