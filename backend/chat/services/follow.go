package services

import (
	"chat/store"
	"chat/types"
)

var _ Follow = (*FollowImpl)(nil)

type FollowImpl struct {
	chatStore *store.ChatStore
}

func (s *FollowImpl) Run() (err error) {
	s.chatStore = store.GetChatStore()
	return
}

func (s *FollowImpl) Stop() {}

func (s *FollowImpl) AddFollow(event *types.RequestEvent) {
	client, ok := s.chatStore.FindClient(event.UserId)
	if !ok {
		return
	}
	if !client.AddFollow(event.TargetUserId) {
		return
	}

	target, ok := s.chatStore.FindClient(event.TargetUserId)
	if !ok || !target.IsActive() {
		return
	}
	client.Write(types.Notification{
		Event:  "follow-login",
		UserId: target.UserId,
	})
}

func (s *FollowImpl) RemoveFollow(event *types.RequestEvent) {
	client, ok := s.chatStore.FindClient(event.UserId)
	if !ok {
		return
	}
	client.RemoveFollow(event.TargetUserId)
}

func (s *FollowImpl) RemoveFollower(event *types.RequestEvent) {
	client, ok := s.chatStore.FindClient(event.UserId)
	if !ok {
		return
	}
	client.RemoveFollower(event.TargetUserId)
}
