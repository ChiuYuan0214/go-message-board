package routes

import (
	"chat/services"
	"chat/types"
	"chat/utils"
	"context"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var _ Handler = (*ChatHandler)(nil)

type ChatHandler struct {
	router       Router
	chatService  services.Chat
	tokenService services.Token
	eventService services.Event
	upgrader     websocket.Upgrader
}

func (h *ChatHandler) Run() (err error) {
	h.upgrader = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool { return true },
	}
	h.router.Handle("/chat", h.handleChats)
	go h.eventService.RunEventLoop()
	return
}

func (h *ChatHandler) Stop() {}

func (h *ChatHandler) handleChats(w http.ResponseWriter, r *http.Request) {
	conn, err := h.upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}

	token := utils.GetTokenFromQuery(r)
	userId := utils.GetUserIdFromToken(token)
	if !h.tokenService.ValidateToken(token, userId) {
		if err := conn.WriteJSON(types.ServerMessage{
			Event:   "error",
			Content: "token invalid.",
		}); err != nil {
			log.Println(err)
		}
		conn.Close()
		return
	}

	h.chatService.InitChatClient(conn, userId, token)
	ctx, cancel := context.WithCancel(context.Background())
	go h.tokenService.UseTokenChecker(ctx, cancel, userId)
	go h.chatService.ListenChatEvent(ctx, cancel, userId)
}
