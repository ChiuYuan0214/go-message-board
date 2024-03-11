package store

import (
	"chat/types"
	"sync"
)

type ChatStore struct {
	Clients   *(map[uint64]*types.Client)
	Broadcast chan *types.RequestEvent
}

func (cs *ChatStore) CreateClient(userId uint64) {
	(*cs.Clients)[userId] = &types.Client{UserId: userId,
		SendMap: &types.SendMap{Lock: sync.Mutex{}, Store: sync.Map{}}}
}

func (cs *ChatStore) GetClient(userId uint64) (*types.Client, bool) {
	_, ok := (*cs.Clients)[userId]
	if !ok {
		cs.CreateClient(userId)
	}
	c := (*cs.Clients)[userId]
	return c, ok
}

func (cs *ChatStore) DeleteClient(userId uint64) {
	delete(*cs.Clients, userId)
}

func (cs *ChatStore) GetSendMap(userId uint64) *types.SendMap {
	client, _ := cs.GetClient(userId)
	return client.SendMap
}

var chatStore ChatStore

func init() {
	chatStore.Clients = &map[uint64]*types.Client{}
	chatStore.Broadcast = make(chan *types.RequestEvent)
}

func GetChatStore() *ChatStore {
	return &chatStore
}
