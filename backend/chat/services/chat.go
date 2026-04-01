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
	clients := s.chatStore.Clients
	client, exist := (*clients)[userId]
	if !exist {
		newClient := &types.Client{
			UserId:   userId,
			Username: "",
			Conn:     conn,
			ConnLock: sync.Mutex{},
			Token:    token,
			IsOnline: true,
			SendMap: &types.SendMap{
				Lock:  sync.Mutex{},
				Store: sync.Map{},
			},
		}
		(*clients)[userId] = newClient
	} else {
		(*client).Conn = conn
		(*client).Token = token
		(*client).IsOnline = true
	}

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
	clients := s.chatStore.Clients
	broadcast := s.chatStore.Broadcast
	client := (*clients)[userId]

	defer func() {
		if _, ok := (*clients)[userId]; ok {
			s.notifyService.NotifyLogout(userId)
		}
		client.Logout()
	}()

	for {
		select {
		case <-ctx.Done():
			return
		default:
			var msg types.RequestEvent
			err := client.Conn.ReadJSON(&msg)
			if err != nil {
				if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
					go s.notifyService.NotifyLogout(userId)
				}
				log.Println(err)
				client.Logout()
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
