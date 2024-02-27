package routes

import (
	"chat/store"
	"chat/types"
	"net/http"

	"github.com/gorilla/websocket"
)

var chatStore = store.GetChatStore()
var cache *types.RedisCache

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func UseCache(redisCache *types.RedisCache) {
	cache = redisCache
}

func UseConnection() {
	http.HandleFunc("/chat", handleChats)
	go handleChatEvents()
}
