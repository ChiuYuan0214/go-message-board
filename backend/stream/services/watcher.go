package services

import (
	"log"
	"stream/types"
)

type WatcherService struct {
	watcher *types.Watcher
	LiveService
}

func NewWatcherService(watcher *types.Watcher) *WatcherService {
	ws := &WatcherService{watcher: watcher}
	ws.SetLive(watcher.GetLive())
	return ws
}

func (ws *WatcherService) Handle() {
	watcher := ws.watcher
	userId := watcher.GetUserId()
	live := watcher.GetLive()
	for {
		messageType, message, err := watcher.GetLiveConn().ReadMessage()
		if err != nil {
			log.Println(err)
			live.WatcherExit(userId)
			return
		}

		switch messageType {
		default:
			log.Println(message)
			return
		}
	}
}
