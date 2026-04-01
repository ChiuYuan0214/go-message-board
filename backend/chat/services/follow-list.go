package services

import (
	"chat/repo"
	"chat/store"
	"chat/types"
	"log"
	"sync"
)

var _ FollowList = (*FollowListImpl)(nil)

type FollowListImpl struct {
	followRepo repo.Follow
	chatStore  *store.ChatStore
}

func (s *FollowListImpl) Run() (err error) {
	s.chatStore = store.GetChatStore()
	return
}

func (s *FollowListImpl) Stop() {}

func (s *FollowListImpl) InitFollowerList(wg *sync.WaitGroup, userId uint64) {
	defer wg.Done()

	clients := s.chatStore.Clients
	followerList, err := s.followRepo.GetFollowerIDs(userId)
	if err != nil {
		log.Println(err)
		return
	}

	onlineList := []uint64{}
	for _, followerId := range followerList {
		follower, ok := (*clients)[followerId]
		if !ok || !follower.IsOnline {
			continue
		}
		onlineList = append(onlineList, followerId)
	}

	client := (*clients)[userId]
	client.FollowerList = followerList
	client.Write(types.UserInfoList{Event: "online-follower-list", List: onlineList})
}

func (s *FollowListImpl) InitFollowList(wg *sync.WaitGroup, userId uint64) {
	defer wg.Done()

	clients := s.chatStore.Clients
	followList, err := s.followRepo.GetFollowIDs(userId)
	if err != nil {
		log.Println(err)
		return
	}

	onlineList := []uint64{}
	for _, followId := range followList {
		follow, ok := (*clients)[followId]
		if !ok || !follow.IsOnline {
			continue
		}
		onlineList = append(onlineList, followId)
	}

	client := (*clients)[userId]
	client.FollowList = followList
	client.Write(types.UserInfoList{Event: "online-follow-list", List: onlineList})
}
