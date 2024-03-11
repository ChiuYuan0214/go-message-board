package jobs

import (
	"chat/types"
	"time"
)

func validateHistoryJob() {
	clients := chatStore.Clients
	go func() {
		for {
			time.Sleep(time.Minute * 10)
			for _, client := range *clients {
				(*client.SendMap).Sync(func() {
					(*client.SendMap).Store.Range(func(key, val any) bool {
						receiverId := key.(uint64)
						msgs := val.([]types.Message)
						newMsgs := []types.Message{}
						for _, msg := range msgs {
							if msg.Ref < 3 || !msg.HasSync {
								newMsgs = append(newMsgs, msg)
							}
						}
						(*client.SendMap).Store.Store(receiverId, newMsgs)
						return true
					})
				})
			}
		}
	}()
}

func incrementHistoryRefJob() {
	clients := chatStore.Clients
	go func() {
		for {
			time.Sleep(time.Minute * 3)
			for _, client := range *clients {
				(*client.SendMap).Sync(func() {
					(*client.SendMap).Store.Range(func(key, val any) bool {
						msgs := val.([]types.Message)
						for _, msg := range msgs {
							msg.Ref = msg.Ref + 1
						}
						return true
					})
				})
			}
		}
	}()
}

func syncHistoryJob() {
	clients := chatStore.Clients
	go func() {
		for {
			time.Sleep(time.Minute * 15)
			for _, c := range *clients {
				client := c
				if client == nil || client.SendMap == nil {
					continue
				}
				go (*client.SendMap).Sync(func() {
					(*client.SendMap).Store.Range(func(key, val any) bool {
						msgs := val.([]types.Message)
						syncList := []interface{}{}
						indexList := []int{}
						for index, msg := range msgs {
							if !msg.HasSync {
								syncList = append(syncList, types.Chat{
									SenderId: msg.UserId, ReceiverId: msg.TargetUserId,
									Content: msg.Content, Time: time.Unix(0, msg.Time)})
								indexList = append(indexList, index)
							}
						}
						if len(syncList) != 0 && mongo.BatchInsert(syncList) {
							for _, index := range indexList {
								msgs[index].HasSync = true
							}
						}
						return true
					})
				})
			}
		}
	}()
}
