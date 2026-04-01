package services

import (
	"chat/repo"
	"chat/store"
	"chat/types"
	"context"
	"log"
	"time"
)

var _ Token = (*TokenImpl)(nil)

type TokenImpl struct {
	tokenRepo repo.Token
	chatStore *store.ChatStore
}

func (s *TokenImpl) Run() (err error) {
	s.chatStore = store.GetChatStore()
	return
}

func (s *TokenImpl) Stop() {}

func (s *TokenImpl) ValidateToken(token string, userId uint64) bool {
	if token == "" || userId == 0 {
		return false
	}

	actualToken, err := s.tokenRepo.GetToken(userId)
	if err != nil {
		log.Println(err)
		return false
	}
	return token == actualToken
}

func (s *TokenImpl) UseTokenChecker(ctx context.Context, cancel context.CancelFunc, userId uint64) {
	client, ok := s.chatStore.FindClient(userId)
	if !ok {
		cancel()
		return
	}

	for {
		time.Sleep(time.Minute * 10)
		select {
		case <-ctx.Done():
			return
		default:
			if !s.ValidateToken(client.TokenValue(), userId) {
				client.Write(types.ServerMessage{
					Event:   "error",
					Content: "token invalid.",
				})
				cancel()
				return
			}
		}
	}
}
