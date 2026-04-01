package service

import (
	"general/repo"
	"general/types"
	"log"
)

var _ Follows = (*FollowsImpl)(nil)

type FollowsImpl struct {
	followerRepo repo.Follower
}

func (s *FollowsImpl) Run() (err error) {
	return
}

func (s *FollowsImpl) Stop() {}

func (s *FollowsImpl) GetFollows(userId uint64) (data []types.Follower) {
	data, err := s.followerRepo.GetUsersByFollowerId(userId)
	if err != nil {
		log.Println(err)
		return []types.Follower{}
	}

	return data
}
