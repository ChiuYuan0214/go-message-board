package services

import (
	"chat/types"
	"context"
	"log"
	"time"
)

func UseTokenChecker(ctx context.Context, cancel context.CancelFunc, userId uint64) {
	client, _ := chatStore.GetClient(userId)
	for {
		time.Sleep(time.Minute * 10)
		select {
		case <-ctx.Done():
			return
		default:
			token := client.Token
			actualToken, err := cache.GetToken(userId)
			if err != nil || token != actualToken {
				log.Println(err)
				client.Write(types.ServerMessage{Event: "error", Content: "token invalid."})
				cancel()
				return
			}
		}
	}
}

func RefreshToken(event *types.RequestEvent) {}
