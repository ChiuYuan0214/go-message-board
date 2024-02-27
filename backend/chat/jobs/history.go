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
					for receiverId, msgs := range (*client.SendMap).Store {
						newMsgs := []types.Message{}
						for _, msg := range msgs {
							if msg.Ref < 3 || !msg.HasSync {
								newMsgs = append(newMsgs, msg)
							}
						}
						(*client.SendMap).Store[receiverId] = newMsgs
					}
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
					for _, msgs := range (*client.SendMap).Store {
						for _, msg := range msgs {
							msg.Ref = msg.Ref + 1
						}
					}
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
				go (*client.SendMap).Sync(func() {
					for receiverId, msgs := range (*client.SendMap).Store {
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
						if len(syncList) == 0 {
							continue
						}
						mongo.BatchInsert(syncList)
						for _, index := range indexList {
							(*client.SendMap).Store[receiverId][index].HasSync = true
						}
					}
				})
			}
		}
	}()
}
