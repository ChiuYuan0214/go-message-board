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
	client, ok := s.chatStore.GetClient(event.UserId)
	if !ok {
		return
	}
	followList := client.FollowList
	for _, id := range followList {
		if id == event.TargetUserId {
			return
		}
	}
	client.FollowList = append(followList, event.TargetUserId)

	target, ok := s.chatStore.GetClient(event.TargetUserId)
	if !ok || !target.IsOnline {
		return
	}
	client.Write(types.Notification{
		Event:  "follow-login",
		UserId: target.UserId,
	})
}

func (s *FollowImpl) RemoveFollow(event *types.RequestEvent) {
	client, ok := s.chatStore.GetClient(event.UserId)
	if !ok {
		return
	}
	var newList []uint64
	for _, id := range client.FollowList {
		if id != event.TargetUserId {
			newList = append(newList, id)
		}
	}
	client.FollowList = newList
}

func (s *FollowImpl) RemoveFollower(event *types.RequestEvent) {
	client, ok := s.chatStore.GetClient(event.UserId)
	if !ok {
		return
	}
	var newList []uint64
	for _, id := range client.FollowerList {
		if id != event.TargetUserId {
			newList = append(newList, id)
		}
	}
	client.FollowerList = newList
}
