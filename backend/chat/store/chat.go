package store

import (
	"chat/types"
)

type ChatStore struct {
	Clients   *(map[uint64]*types.Client)
	Broadcast chan *types.RequestEvent
}

func (cs *ChatStore) GetClient(userId uint64) (*types.Client, bool) {
	c, ok := (*cs.Clients)[userId]
	return c, ok
}

func (cs *ChatStore) DeleteClient(userId uint64) {
	delete(*cs.Clients, userId)
}

func (cs *ChatStore) GetSendMap(userId uint64) *types.SendMap {
	store, ok := cs.GetClient(userId)
	if !ok {
		return &types.SendMap{}
	}
	return store.SendMap
}

var chatStore ChatStore

func init() {
	chatStore.Clients = &map[uint64]*types.Client{}
	chatStore.Broadcast = make(chan *types.RequestEvent)
}

func GetChatStore() *ChatStore {
	return &chatStore
}
