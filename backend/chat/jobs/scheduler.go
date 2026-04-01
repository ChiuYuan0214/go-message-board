package jobs

import (
	"chat/repo"
	"chat/store"
	"chat/types"
	"time"
)

var _ Scheduler = (*SchedulerImpl)(nil)

type SchedulerImpl struct {
	historyRepo repo.History
	chatStore   *store.ChatStore
}

func (s *SchedulerImpl) Run() (err error) {
	s.chatStore = store.GetChatStore()
	s.validateHistoryJob()
	s.incrementHistoryRefJob()
	s.syncHistoryJob()
	s.removeLogoutUsersJob()
	return
}

func (s *SchedulerImpl) Stop() {}

func (s *SchedulerImpl) validateHistoryJob() {
	clients := s.chatStore.Clients
	go func() {
		for {
			time.Sleep(time.Minute * 10)
			for _, client := range *clients {
				(*client.SendMap).Sync(func() {
					sendMap := client.SendMap
					sendMap.Store.Range(func(key, val any) bool {
						receiverId := key.(uint64)
						msgs := val.([]types.Message)
						var newMsgs []types.Message
						for _, msg := range msgs {
							if sendMap.MapRef < 3 || !msg.HasSync {
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

func (s *SchedulerImpl) incrementHistoryRefJob() {
	clients := s.chatStore.Clients
	go func() {
		for {
			time.Sleep(time.Minute * 3)
			for _, client := range *clients {
				(*client.SendMap).Sync(func() {
					(*client.SendMap).MapRef++
				})
			}
		}
	}()
}

func (s *SchedulerImpl) syncHistoryJob() {
	clients := s.chatStore.Clients
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
						var syncList []types.DynamoChat
						var indexList []int
						for index, msg := range msgs {
							if !msg.HasSync {
								syncList = append(syncList, types.DynamoChat{
									SenderId:   msg.UserId,
									ReceiverId: msg.TargetUserId,
									Content:    msg.Content,
									Time:       time.Unix(0, msg.Time),
								})
								indexList = append(indexList, index)
							}
						}
						if len(syncList) != 0 && s.historyRepo.BatchInsert(syncList) {
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

func (s *SchedulerImpl) removeLogoutUsersJob() {
	go func() {
		for {
			time.Sleep(30 * time.Minute)
			for userId, client := range *s.chatStore.Clients {
				if !client.IsOnline && client.LogoutTime.Before(time.Now().Add(-10*time.Minute)) {
					s.chatStore.DeleteClient(userId)
				}
			}
		}
	}()
}
