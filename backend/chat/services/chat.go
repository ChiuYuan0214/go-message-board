package services

import (
	"chat/store"
	"chat/types"
	"context"
	"log"
	"sync"

	"github.com/gorilla/websocket"
)

var _ Chat = (*ChatImpl)(nil)

type ChatImpl struct {
	followListService FollowList
	notifyService     Notify
	chatStore         *store.ChatStore
}

func (s *ChatImpl) Run() (err error) {
	s.chatStore = store.GetChatStore()
	return
}

func (s *ChatImpl) Stop() {}

func (s *ChatImpl) InitChatClient(conn *websocket.Conn, userId uint64, token string) {
	client, exist := s.chatStore.GetClient(userId)
	if !exist {
		client.Username = ""
		client.SendMap = &types.SendMap{
			Lock:  sync.Mutex{},
			Store: sync.Map{},
		}
	}
	client.InitSession(conn, token)

	go func() {
		var wg sync.WaitGroup
		wg.Add(1)
		go s.followListService.InitFollowerList(&wg, userId)
		wg.Add(1)
		go s.followListService.InitFollowList(&wg, userId)
		wg.Wait()
		s.notifyService.NotifyLogin(userId)
	}()
}

func (s *ChatImpl) ListenChatEvent(ctx context.Context, cancel context.CancelFunc, userId uint64) {
	broadcast := s.chatStore.Broadcast
	client, ok := s.chatStore.FindClient(userId)
	if !ok {
		cancel()
		return
	}

	defer func() {
		if _, ok := s.chatStore.FindClient(userId); ok {
			s.notifyService.NotifyLogout(userId)
		}
		client.CloseSession()
	}()

	for {
		select {
		case <-ctx.Done():
			return
		default:
			var msg types.RequestEvent
			err := client.Read(&msg)
			if err != nil {
				log.Println(err)
				client.CloseSession()
				cancel()
				return
			}
			if msg.Type == "ping" {
				continue
			}
			if msg.UserId != userId {
				client.Write(types.ServerMessage{
					Event:   "error",
					Content: "userId incorrect.",
				})
				continue
			}
			broadcast <- &msg
		}
	}
}
