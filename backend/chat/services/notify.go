package services

import (
	"chat/store"
	"chat/types"
)

var _ Notify = (*NotifyImpl)(nil)

type NotifyImpl struct {
	chatStore *store.ChatStore
}

func (s *NotifyImpl) Run() (err error) {
	s.chatStore = store.GetChatStore()
	return
}

func (s *NotifyImpl) Stop() {}

func (s *NotifyImpl) NotifyLogin(userId uint64) {
	client, exist := s.chatStore.FindClient(userId)
	if !exist {
		return
	}

	for _, id := range client.SnapshotFollowerList() {
		follower, ok := s.chatStore.FindClient(id)
		if ok && follower.IsActive() {
			go follower.Write(types.Notification{
				Event:  "follow-login",
				UserId: userId,
			})
		}
	}

	for _, id := range client.SnapshotFollowList() {
		follow, ok := s.chatStore.FindClient(id)
		if ok && follow.IsActive() {
			go follow.Write(types.Notification{
				Event:  "follower-login",
				UserId: userId,
			})
		}
	}
}

func (s *NotifyImpl) NotifyLogout(userId uint64) {
	client, exist := s.chatStore.FindClient(userId)
	if !exist {
		return
	}

	for _, id := range client.SnapshotFollowerList() {
		follower, ok := s.chatStore.FindClient(id)
		if ok && follower.IsActive() {
			go follower.Write(types.Notification{
				Event:  "follow-logout",
				UserId: userId,
			})
		}
	}

	for _, id := range client.SnapshotFollowList() {
		follow, ok := s.chatStore.FindClient(id)
		if ok && follow.IsActive() {
			go follow.Write(types.Notification{
				Event:  "follower-logout",
				UserId: userId,
			})
		}
	}
}
