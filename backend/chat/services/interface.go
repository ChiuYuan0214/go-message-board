package services

import (
	"chat/types"
	"context"
	"sync"

	"github.com/gorilla/websocket"
)

type Chat interface {
	InitChatClient(conn *websocket.Conn, userId uint64, token string)
	ListenChatEvent(ctx context.Context, cancel context.CancelFunc, userId uint64)
}

type Token interface {
	ValidateToken(token string, userId uint64) bool
	UseTokenChecker(ctx context.Context, cancel context.CancelFunc, userId uint64)
}

type Event interface {
	RunEventLoop()
	HandleEvent(event *types.RequestEvent)
}

type History interface {
	GetHistory(event *types.RequestEvent)
}

type Message interface {
	SendMessage(reqMsg *types.RequestEvent)
}

type Follow interface {
	AddFollow(event *types.RequestEvent)
	RemoveFollow(event *types.RequestEvent)
	RemoveFollower(event *types.RequestEvent)
}

type FollowList interface {
	InitFollowerList(wg *sync.WaitGroup, userId uint64)
	InitFollowList(wg *sync.WaitGroup, userId uint64)
}

type Notify interface {
	NotifyLogin(userId uint64)
	NotifyLogout(userId uint64)
}
