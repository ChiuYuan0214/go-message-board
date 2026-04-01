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
	client, exist := s.chatStore.GetClient(userId)
	if !exist {
		return
	}

	for _, id := range client.FollowerList {
		follower, ok := s.chatStore.GetClient(id)
		if ok && follower.IsOnline {
			go follower.Write(types.Notification{Event: "follow-login", UserId: userId})
		}
	}

	for _, id := range client.FollowList {
		follow, ok := s.chatStore.GetClient(id)
		if ok && follow.IsOnline {
			go follow.Write(types.Notification{Event: "follower-login", UserId: userId})
		}
	}
}

func (s *NotifyImpl) NotifyLogout(userId uint64) {
	client, exist := s.chatStore.GetClient(userId)
	if !exist {
		return
	}

	for _, id := range client.FollowerList {
		follower, ok := s.chatStore.GetClient(id)
		if ok && follower.IsOnline {
			go follower.Write(types.Notification{Event: "follow-logout", UserId: userId})
		}
	}

	for _, id := range client.FollowList {
		follow, ok := s.chatStore.GetClient(id)
		if ok && follow.IsOnline {
			go follow.Write(types.Notification{Event: "follower-logout", UserId: userId})
		}
	}
}
