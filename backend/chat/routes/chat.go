package routes

import (
	"chat/services"
	"chat/store"
	"chat/types"
	"chat/utils"
	"context"
	"log"
	"net/http"
)

func handleChats(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	token := utils.GetTokenFromQuery(r)
	userId := utils.GetUserIdFromToken(token)
	actualToken, err := cache.GetToken(userId)
	if err != nil || token != actualToken {
		log.Println(err)
		conn.WriteJSON(types.ServerMessage{Event: "error", Content: "token invalid."})
		return
	}

	services.InitChatClient(conn, userId, token)
	ctx, cancel := context.WithCancel(context.Background())
	go services.UseTokenChecker(ctx, cancel, userId)
	go services.ListenChatEvent(ctx, cancel, userId)
}

func handleChatEvents() {
	broadcast := store.GetChatStore().Broadcast
	for {
		event := <-broadcast

		switch event.Type {
		case "message":
			go services.SendMessage(event)
		case "add-follow":
			go services.AddFollow(event)
		case "remove-follow":
			go services.RemoveFollow(event)
		case "remove-follower":
			go services.RemoveFollower(event)
		case "refresh-token":
			go services.RefreshToken(event)
		case "history":
			go services.GetHistory(event)
		}
	}
}
