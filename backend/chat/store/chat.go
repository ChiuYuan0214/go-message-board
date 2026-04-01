package store

import (
	"chat/types"
	"sync"
)

type ChatStore struct {
	mu        sync.RWMutex
	clients   map[uint64]*types.Client
	Broadcast chan *types.RequestEvent
}

func (cs *ChatStore) CreateClient(userId uint64) *types.Client {
	client := &types.Client{
		StateLock: sync.RWMutex{},
		UserId:    userId,
		SendMap:   new(types.SendMap),
	}
	cs.clients[userId] = client
	return client
}

func (cs *ChatStore) GetClient(userId uint64) (*types.Client, bool) {
	cs.mu.Lock()
	defer cs.mu.Unlock()

	client, ok := cs.clients[userId]
	if !ok {
		client = cs.CreateClient(userId)
	}
	return client, ok
}

func (cs *ChatStore) FindClient(userId uint64) (*types.Client, bool) {
	cs.mu.RLock()
	defer cs.mu.RUnlock()

	client, ok := cs.clients[userId]
	return client, ok
}

func (cs *ChatStore) SnapshotClients() map[uint64]*types.Client {
	cs.mu.RLock()
	defer cs.mu.RUnlock()

	snapshot := make(map[uint64]*types.Client, len(cs.clients))
	for userId, client := range cs.clients {
		snapshot[userId] = client
	}
	return snapshot
}

func (cs *ChatStore) DeleteClient(userId uint64) {
	cs.mu.Lock()
	defer cs.mu.Unlock()

	delete(cs.clients, userId)
}

func (cs *ChatStore) GetSendMap(userId uint64) *types.SendMap {
	client, _ := cs.GetClient(userId)
	return client.SendMap
}

var chatStore ChatStore

func init() {
	chatStore.clients = map[uint64]*types.Client{}
	chatStore.Broadcast = make(chan *types.RequestEvent)
}

func GetChatStore() *ChatStore {
	return &chatStore
}
