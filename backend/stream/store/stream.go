package store

import (
	"stream/types"
	"sync"

	"github.com/gorilla/websocket"
)

var ownerMap = map[uint64]*types.Live{}

var lives = &sync.Map{}

func GetOwnerMap() map[uint64]*types.Live {
	return ownerMap
}

func GetLives() *sync.Map {
	return lives
}

func CreateOwner(liveId, userId uint64, conn *websocket.Conn) *types.Owner {
	if _, ok := lives.Load(liveId); !ok {
		lives.Store(liveId, types.InitLive(liveId))
	}
	val, _ := lives.Load(liveId)
	live := val.(*types.Live)
	owner := live.SetOwnerLiveConn(conn)
	lives.Store(liveId, live)
	ownerMap[userId] = live
	return owner
}

func CreateWatcher(liveId, userId uint64, conn *websocket.Conn) *types.Watcher {
	if _, ok := lives.Load(liveId); !ok {
		lives.Store(liveId, types.InitLive(liveId))
	}
	val, _ := lives.Load(liveId)
	live := val.(*types.Live)
	watcher := live.SetWatcherLiveConn(userId, conn)
	return watcher
}

func CreateOwnerRecord(liveId, userId uint64, conn *websocket.Conn) *types.Owner {
	if _, ok := lives.Load(liveId); !ok {
		lives.Store(liveId, types.InitLive(liveId))
	}
	val, _ := lives.Load(liveId)
	live := val.(*types.Live)
	owner := live.SetOwnerRecordConn(conn)
	return owner
}

func CreateWatcherRecord(liveId, userId uint64, conn *websocket.Conn) *types.Watcher {
	if _, ok := lives.Load(liveId); !ok {
		lives.Store(liveId, types.InitLive(liveId))
	}
	val, _ := lives.Load(liveId)
	live := val.(*types.Live)
	watcher := live.SetWatcherRecordConn(userId, conn)
	return watcher
}
