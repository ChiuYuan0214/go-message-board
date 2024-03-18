package services

import (
	"log"
	"stream/types"
)

type LiveService struct {
	live *types.Live
}

func (c *LiveService) SetLive(live *types.Live) {
	c.live = live
}

func (c *LiveService) PushStream(senderId uint64, watchers *types.WatcherMap, message []byte) {
	watchers.Range(func(userId, watcherClient any) bool {
		id := userId.(uint64)
		c := watcherClient.(*types.Client)
		if id == senderId {
			return true
		}
		go func() {
			err := c.Write(message)
			if err != nil {
				log.Println(err)
				c.Close()
				c.GetLive().WatcherExit(id)
			}
		}()
		return true
	})
}
