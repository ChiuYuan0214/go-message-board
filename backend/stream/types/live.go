package types

import (
	"sync"

	"github.com/gorilla/websocket"
)

type WatcherMap = sync.Map

type Live struct {
	liveId         uint64
	owner          *Owner
	watchers       *WatcherMap
	watcherCntLock sync.RWMutex
	watcherCount   uint
	isStart        bool
}

func InitLive(liveId uint64) *Live {
	return &Live{liveId: liveId, owner: &Owner{}, watchers: &WatcherMap{}, isStart: true}
}

func (live *Live) GetLiveId() uint64 {
	return live.liveId
}

func (live *Live) GetOwner() *Owner {
	return live.owner
}

func (live *Live) OwnerExit() {
	live.isStart = false
}

func (live *Live) WatcherJoin(userId uint64, client *Watcher) {
	live.watchers.Store(userId, client)
	live.watcherCntLock.Lock()
	live.watcherCount++
	live.watcherCntLock.Unlock()
}

func (live *Live) WatcherExit(userId uint64) {
	live.watchers.Delete(userId)
	live.watcherCntLock.Lock()
	live.watcherCount--
	live.watcherCntLock.Unlock()
}

func (live *Live) SetOwnerLiveConn(conn *websocket.Conn) *Owner {
	owner := live.GetOwner()
	if owner == nil {
		owner = &Owner{}
	}
	owner.SetLiveConn(conn)
	owner.SetLive(live)
	live.owner = owner
	return owner
}

func (live *Live) SetOwnerRecordConn(conn *websocket.Conn) *Owner {
	owner := live.GetOwner()
	if owner == nil {
		owner = &Owner{}
	}
	owner.SetRecordConn(conn)
	owner.SetLive(live)
	live.owner = owner
	return owner
}

func (live *Live) SetWatcherLiveConn(userId uint64, conn *websocket.Conn) *Watcher {
	watcher := live.GetWatcher(userId)
	if watcher == nil {
		watcher = &Watcher{}
	}
	watcher.SetLiveConn(conn)
	watcher.SetLive(live)
	live.WatcherJoin(userId, watcher)
	return watcher
}

func (live *Live) SetWatcherRecordConn(userId uint64, conn *websocket.Conn) *Watcher {
	watcher := live.GetWatcher(userId)
	if watcher == nil {
		watcher = &Watcher{}
	}
	watcher.SetRecordConn(conn)
	watcher.SetLive(live)
	live.WatcherJoin(userId, watcher)
	return watcher
}

func (live *Live) HasWatcher(userId uint64) bool {
	_, ok := live.watchers.Load(userId)
	return ok
}

func (live *Live) GetWatchers() *WatcherMap {
	return live.watchers
}

func (live *Live) GetWatcher(userId uint64) *Watcher {
	val, ok := live.watchers.Load(userId)
	if !ok {
		return nil
	}
	return val.(*Watcher)
}
