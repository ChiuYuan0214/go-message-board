package services

import (
	"chat/types"
	"context"
	"log"
	"sync"

	"github.com/gorilla/websocket"
)

func InitChatClient(conn *websocket.Conn, userId uint64, token string) {
	clients := chatStore.Clients
	client, exist := (*clients)[userId]
	if !exist {
		newClient := &types.Client{UserId: userId, Username: "", Conn: conn, ConnLock: sync.Mutex{},
			Token: token, IsOnline: true, SendMap: &types.SendMap{Lock: sync.Mutex{}, Store: map[uint64][]types.Message{}}}
		(*clients)[userId] = newClient
	} else {
		(*client).Conn = conn
		(*client).Token = token
		(*client).IsOnline = true
	}

	go func() {
		var wg sync.WaitGroup
		wg.Add(1)
		go InitFollowerList(&wg, conn, userId)
		wg.Add(1)
		go InitFollowList(&wg, conn, userId)
		wg.Wait()
		NotifyLogin(userId)
	}()
}

func ListenChatEvent(ctx context.Context, cancel context.CancelFunc, userId uint64) {
	clients := chatStore.Clients
	broadcast := chatStore.Broadcast
	client := (*clients)[userId]

	defer func() {
		if _, ok := (*clients)[userId]; ok {
			NotifyLogout(userId)
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
					go NotifyLogout(userId)
				}
				log.Println(err)
				client.Logout()
				cancel()
				return
			}
			if msg.UserId != userId {
				client.Write(types.ServerMessage{Event: "error", Content: "userId incorrect."})
				continue
			}
			broadcast <- &msg
		}
	}
}
