package services

import (
	"log"
	"stream/types"
)

type LiveRecordService struct {
	live *types.Live
}

func (lrs *LiveRecordService) SetLive(live *types.Live) {
	lrs.live = live
}

func (lrs *LiveRecordService) SendChatOrReaction(senderId uint64, message []byte) {
	watchers := lrs.live.GetWatchers()
	watchers.Range(func(key, val any) bool {
		id := key.(uint64)
		otherWatcher := val.(*types.Watcher)
		if id == senderId {
			return true
		}
		go func() {
			err := otherWatcher.Write(message)
			if err != nil {
				log.Println(err)
				otherWatcher.Close()
				lrs.live.WatcherExit(id)
			}
		}()
		return true
	})
}
